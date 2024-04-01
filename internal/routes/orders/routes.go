package orders

import (
	"shop/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(app *fiber.App, handler handler.OrderHandler){
	app.Post("/order", handler.AddOrder)
	app.Get("/order/:session_key", handler.GetOrder)
}