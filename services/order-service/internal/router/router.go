package router

import (
	"order-service/internal/handler"

	"github.com/gofiber/fiber/v2"
)


func SetupRoutes(app *fiber.App,h *handler.OrderHandler) {

	app.Post("/order",h.CreateOrder)
	app.Get("/orders",h.GetAllOrders)
	app.Get("/orders/:id",h.GetOrderByID)
	

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})
}