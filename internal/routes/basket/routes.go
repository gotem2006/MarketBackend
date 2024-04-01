package basket

import (
	"shop/internal/handler"
	"github.com/gofiber/fiber/v2"
)



func BasketRoutes(app *fiber.App, handler handler.BasketHandler){
	app.Post("/basket", handler.AddBasket)
	app.Get("/basket", handler.GetBasket)
	app.Post("/basket/quantity", handler.ReduceQuantity)
	app.Post("/basket/delete", handler.DeleteBasket)
}