package orders

import (
	"shop/internal/service"

	"github.com/gofiber/fiber/v2"
)



type orderHandler struct{
	orderService service.OrderService
}


func NewOrderHandler(orderService service.OrderService) *orderHandler{
	return &orderHandler{orderService}
}



func (h *orderHandler) AddOrder(c *fiber.Ctx) error{
	return h.orderService.AddOrder(c)
}



func (h *orderHandler) GetOrder(c *fiber.Ctx) error{
	order, err := h.orderService.GetOrder(c)
	if err != nil{
		return c.JSON(err)
	}
	return c.JSON(order)
}