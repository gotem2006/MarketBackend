package service

import (
	model "shop/internal/models"
	"shop/internal/repository"
	"shop/internal/service/basket"
	"shop/internal/service/orders"
	"shop/internal/service/product"

	"github.com/gofiber/fiber/v2"
)


type ProductService interface{
	GetProduct(c *fiber.Ctx) (Product model.Product, err error)
	GetProducts(c *fiber.Ctx) ([]model.Product, error)
	AddProduct(c *fiber.Ctx) error
	AddAttribute(c *fiber.Ctx) error
	AddCategory(c *fiber.Ctx) error
}

type BasketService interface{
	AddBasket(c *fiber.Ctx) error
	GetBasket(c *fiber.Ctx) (Baskets []model.Basket, err error)
	ReduceQuantity(c *fiber.Ctx) error
	DeleteBasket(c *fiber.Ctx) error
}

type OrderService interface{
	AddOrder(c *fiber.Ctx) error
	GetOrder(c *fiber.Ctx) (model.Order, error)
}


type Services struct{
	Product ProductService
	Basket BasketService
	Order OrderService
}

func NewServices(repo *repository.Repositories) *Services{
	productService := product.NewProductService(repo.Product)
	basketService := basket.NewBasketService(repo.Basket)
	orderService := orders.NewOrderService(repo.Order)
	return &Services{
		Product: productService,
		Basket: basketService,
		Order: orderService,
	}
}