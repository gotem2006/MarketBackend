package basket

import (
	"shop/internal/handler"
	"github.com/gofiber/fiber/v2"
)



func BasketRoutes(app *fiber.App, handler handler.BasketHandler){
	app.Post("add/:id", func(c *fiber.Ctx) error {return c.JSON(fiber.Map{"response": "pashel naxui"})})
}