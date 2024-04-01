package basket

import (
	"log"
	model "shop/internal/models"
	"shop/internal/repository"

	"github.com/gofiber/fiber/v2"
)





type basketService struct{
	repo repository.BasketRepository
}
func NewBasketService(repo repository.BasketRepository)*basketService{
	return &basketService{repo: repo}
}



func (b *basketService) AddBasket(c *fiber.Ctx) error{
	basket := new(model.Basket)
	if err := c.BodyParser(&basket); err != nil{
		log.Println(err)
	}
	return b.repo.AddBasket(*basket)
}



func (b *basketService) GetBasket(c *fiber.Ctx) (Baskets []model.Basket, err error){
	var session_key struct{
		SessionKey string `json:"session_key"`
	}
	if err := c.BodyParser(&session_key); err != nil{
		return nil, err
	}
	Baskets, err = b.repo.GetBasket(session_key.SessionKey)
	if err != nil{
		return nil, err
	}
	return Baskets, nil
}

func (b *basketService) ReduceQuantity(c *fiber.Ctx) error{
	basket := new(model.Basket)
	if err := c.BodyParser(&basket); err != nil{
		return err
	}
	return b.repo.ReduceQuantity(*basket)
}


func (b *basketService) DeleteBasket(c *fiber.Ctx) error{
	basket := new(model.Basket)
	if err := c.BodyParser(&basket); err != nil{
		return err
	}
	return b.repo.DeleteBasket(*basket)
}