package routes

import (
	"shop/internal/handler"
	"shop/internal/routes/basket"
	product "shop/internal/routes/products"
	"shop/internal/routes/orders"
	"github.com/gofiber/fiber/v2"
)


func InitRoutes(app *fiber.App, handlers handler.Handlers){
	basket.BasketRoutes(app, handlers.Basket)
	product.ProductRoutes(app, handlers.Product)
	orders.OrderRoutes(app, handlers.Order)
}