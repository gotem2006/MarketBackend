package product

import (
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
	response, err := p.productService.GetProduct(c)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}
	return c.JSON(response)
}


func (p *productHandler) PostProduct(c *fiber.Ctx) error{
	err := p.productService.AddProduct(c)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}
	return c.JSON(fiber.Map{"Result": "Succes, product added"})
}

func (p *productHandler) GetProducts(c *fiber.Ctx) error{
	response, err := p.productService.GetProducts(c)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}
	return c.JSON(fiber.Map{"products": response})
}

func (p *productHandler) AddAttribute(c *fiber.Ctx) error{
	err := p.productService.AddAttribute(c)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}
	return c.JSON(fiber.Map{"Result": "Success, attribute added"})
}

func (p *productHandler) AddCategory(c *fiber.Ctx) error{
	err := p.productService.AddCategory(c)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}
	return c.JSON(fiber.Map{"Result": "Success, category added"})
}