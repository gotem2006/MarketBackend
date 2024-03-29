package service

import (
	model "shop/internal/models"
	"shop/internal/repository"
	"shop/internal/service/basket"
	"shop/internal/service/product"

	"github.com/gofiber/fiber/v2"
)


type ProductService interface{
	GetProduct(c *fiber.Ctx) model.Product
	GetProducts(c *fiber.Ctx) []model.Product
	AddProduct(c *fiber.Ctx) error
	AddAttribute(c *fiber.Ctx) error
	AddCategory(c *fiber.Ctx) error
}

type BasketService interface{

}



type Services struct{
	Product ProductService
	Basket BasketService
}

func NewServices(repo *repository.Repositories) *Services{
	productService := product.NewProductService(repo.Product)
	basketService := basket.NewBasketService(repo)
	return &Services{
		Product: productService,
		Basket: basketService,
	}
}