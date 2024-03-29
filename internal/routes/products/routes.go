package product

import (
	"shop/internal/handler"
	"github.com/gofiber/fiber/v2"
)


func ProductRoutes(app *fiber.App, handler handler.ProductHandler){
	app.Get("/products", handler.GetProducts)
	app.Get("/product/:id", handler.GetProduct)
	app.Post("/attribute/:id", handler.AddAttribute)
	app.Post("/product/", handler.PostProduct)
	app.Post("/category", handler.AddCategory)
}