package orders

import (
	model "shop/internal/models"
	"shop/internal/repository"

	"github.com/gofiber/fiber/v2"
)


type orderService struct{
	orderRepository repository.OrderRepository
}



func NewOrderService(orderRepository repository.OrderRepository) *orderService{
	return &orderService{orderRepository}
}


func (o *orderService) AddOrder(c *fiber.Ctx) error{
	var order model.Order
	if err := c.BodyParser(&order); err != nil{
		return err
	}
	return o.orderRepository.AddOrder(order)
}



func (o *orderService) GetOrder(c *fiber.Ctx) (model.Order, error){
	sessionKey := c.Params("session_key")
	return o.orderRepository.GetOrder(sessionKey)
}