package main

import (
	"fmt"
	"log"

	"auth-service/internal/config"
	"auth-service/internal/db"
	"auth-service/internal/handler"
	"auth-service/internal/repository"
	"auth-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	database := db.Connect(dbURL)

	// layers
	repo := repository.NewUserRepository(database)
	jwtSvc := service.NewJWTService(cfg.JWTSecret)
	svc := service.NewUserService(repo,jwtSvc)
	h := handler.NewUserHandler(svc)

	app := fiber.New()

	app.Post("/register", h.Register)
	app.Post("/login", h.Login)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	log.Println("Auth service running on port", cfg.AppPort)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}