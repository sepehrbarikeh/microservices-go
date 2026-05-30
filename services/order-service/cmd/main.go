package main

import (
	"context"
	"fmt"
	"log"

	"order-service/internal/config"
	"order-service/internal/db"
	"order-service/internal/handler"
	"order-service/internal/repository"
	"order-service/internal/router"
	"order-service/internal/service"

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

	database.Ping(context.Background())
	// layers
	repo := repository.NewOrderRepository(database)
	svc := service.NewOrderService(repo)
	handler := handler.NewOrderHandler(svc)

	app := fiber.New()

	router.SetupRoutes(app,handler)


	log.Println("Auth service running on port", cfg.AppPort)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
