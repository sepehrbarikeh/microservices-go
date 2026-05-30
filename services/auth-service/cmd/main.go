package main

import (
	"fmt"
	"log"

	"auth-service/internal/config"
	"auth-service/internal/db"
	"auth-service/internal/handler"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"auth-service/router"

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
	svc := service.NewUserService(repo, jwtSvc)
	handler := handler.NewUserHandler(svc)

	app := fiber.New()

	router.SetupRoutes(app, handler, cfg.JWTSecret)

	log.Println("Auth service running on port", cfg.AppPort)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
