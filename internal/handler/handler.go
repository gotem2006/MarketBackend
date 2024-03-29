package handler

import (
	"shop/internal/handler/basket"
	"shop/internal/handler/product"
	"shop/internal/service"
	"github.com/gofiber/fiber/v2"
)



type ProductHandler interface{
	GetProduct(c *fiber.Ctx) error
	PostProduct(c *fiber.Ctx) error
	GetProducts(c *fiber.Ctx) error
	AddAttribute(c *fiber.Ctx) error
	AddCategory(c *fiber.Ctx) error
}

type BasketHandler interface{
	
}


type Handlers struct{
	Product ProductHandler
	Basket BasketHandler
}


func NewHandlers(services service.Services) *Handlers{
	productHandler := product.NewProductHandler(services.Product)
	basketHandler := basket.NewBasketHandler(services.Basket)
	return &Handlers{
		Product: productHandler,
		Basket: basketHandler,
	}
}


