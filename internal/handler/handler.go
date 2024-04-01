package handler

import (
	"shop/internal/handler/basket"
	"shop/internal/handler/orders"
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
	AddBasket(c *fiber.Ctx) error
	GetBasket(c *fiber.Ctx) error
	ReduceQuantity(c *fiber.Ctx) error
	DeleteBasket(c *fiber.Ctx) error
}


type OrderHandler interface{
	AddOrder(c *fiber.Ctx) error
	GetOrder(c *fiber.Ctx) error
}


type Handlers struct{
	Product ProductHandler
	Basket BasketHandler
	Order OrderHandler
}


func NewHandlers(services service.Services) *Handlers{
	productHandler := product.NewProductHandler(services.Product)
	basketHandler := basket.NewBasketHandler(services.Basket)
	orderHandler := orders.NewOrderHandler(services.Order)
	return &Handlers{
		Product: productHandler,
		Basket: basketHandler,
		Order: orderHandler,
	}
}


