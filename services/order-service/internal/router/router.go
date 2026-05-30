package router

import (
	"order-service/internal/handler"
	"order-service/internal/middleware"

	"github.com/gofiber/fiber/v2"
)


func SetupRoutes(app *fiber.App,h *handler.OrderHandler,secretKey string) {

	app.Group("/",middleware.JWTMiddleware(secretKey))

	app.Post("/order",h.CreateOrder)
	app.Get("/orders",h.GetAllOrders)
	app.Get("/orders/:id",h.GetOrderByID)
	

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})
}