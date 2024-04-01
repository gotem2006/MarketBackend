package basket

import (
	"shop/internal/service"

	"github.com/gofiber/fiber/v2"
)





type basketHandler struct{
	BasketServcie service.BasketService
}



func NewBasketHandler(basketService service.BasketService) *basketHandler{
	return &basketHandler{
		BasketServcie: basketService,
	}
}





func (b *basketHandler) AddBasket(c *fiber.Ctx) error{
	err := b.BasketServcie.AddBasket(c)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}
	return c.JSON(fiber.Map{"Result": "Success, product added to basket"})
}


func (b *basketHandler) GetBasket(c *fiber.Ctx) error{
	response, err := b.BasketServcie.GetBasket(c)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}
	return c.JSON(response)
}


func (b *basketHandler) ReduceQuantity(c *fiber.Ctx) error{
	err := b.BasketServcie.ReduceQuantity(c)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}
	return c.JSON(fiber.Map{"Result": "Success, quantity reduced"})
}

func (b *basketHandler) DeleteBasket(c *fiber.Ctx) error{
	err := b.BasketServcie.DeleteBasket(c)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err})
	}
	return c.JSON(fiber.Map{"Result": "Success, products from basket deleted"})
}