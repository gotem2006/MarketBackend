package product

import (
	model "shop/internal/models"
	"shop/internal/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)



type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *productService{
	return &productService{repo: repo}
}


func (p *productService) GetProduct(c *fiber.Ctx) (Product model.Product, err error){
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil{
		return Product, err
	}
	Product, err = p.repo.GetProduct(id)
	if err != nil{
		return Product, err
	}
	return Product, nil
}

func (p *productService) GetProducts(c *fiber.Ctx) ([]model.Product, error){
	return p.repo.GetProducts()
}


func (p *productService) AddProduct(c *fiber.Ctx) error{
	product := new(model.Product)
	if err := c.BodyParser(&product);err != nil{
		return err
	}
	return p.repo.AddProduct(*product)
}

func (p *productService) AddAttribute(c *fiber.Ctx) error{
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil{
		return err
	}
	attribute := new(model.Attribute)
	if err := c.BodyParser(&attribute); err != nil{
		return err
	} 
	return p.repo.AddAttribute(*attribute, id)
}


func (p *productService) AddCategory(c *fiber.Ctx) error{
	category := new(model.Category)
	if err := c.BodyParser(&category); err != nil{
		return err
	}
	return p.repo.AddCategory(*category)
}

