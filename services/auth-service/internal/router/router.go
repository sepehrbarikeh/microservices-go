package router

import (
	"auth-service/internal/handler"
	"auth-service/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h *handler.UserHandler, jwtSecret string) {
	app.Post("/register", h.Register)
	app.Post("/login", h.Login)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	api := app.Group("/")
	api.Use(middleware.JWTMiddleware(jwtSecret))
	api.Get("/profile", h.Profile)
}
