package handler

import (
	model "order-service/internal/entity"
	"order-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order model.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	err := h.service.CreateOrder(order)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "order created",
	})

}

func (h *OrderHandler) GetOrderByID(c *fiber.Ctx) error {
	OrderID := c.Params("id")

	order,err := h.service.GetOrderByID(OrderID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(order)

}

func (h *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	
	order,err := h.service.GetAllOrders()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(order)

}

