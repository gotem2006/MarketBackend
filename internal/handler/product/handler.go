package product

import (
	"log"
	"shop/internal/service"

	"github.com/gofiber/fiber/v2"
)



type productHandler struct{
	productService service.ProductService
}


func NewProductHandler(productService service.ProductService) *productHandler{
	return &productHandler{productService: productService}
}


func (p *productHandler) GetProduct(c *fiber.Ctx) error{
	return c.JSON(p.productService.GetProduct(c))
}


func (p *productHandler) PostProduct(c *fiber.Ctx) error{
	err := p.productService.AddProduct(c)
	if err != nil{
		log.Println(err)
		return c.JSON(fiber.Map{"result": "bad request"})
	}
	return c.JSON(fiber.Map{"result": "succes"})
}

func (p *productHandler) GetProducts(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{"products": p.productService.GetProducts(c)})
}

func (p *productHandler) AddAttribute(c *fiber.Ctx) error{
	return p.productService.AddAttribute(c)
}

func (p *productHandler) AddCategory(c *fiber.Ctx) error{
	return p.productService.AddCategory(c)
}