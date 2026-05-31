package main

import (
	"context"
	"fmt"
	"log"

	"order-service/internal/config"
	"order-service/internal/db"
	"order-service/internal/grpc/client"
	"order-service/internal/handler"
	"order-service/internal/repository"
	"order-service/internal/router"
	"order-service/internal/service"
	"order-service/rabbitmq"

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
	rabbitmq, err := rabbitmq.NewRabbitMQ(cfg.MQUser, cfg.MQPassword, cfg.MQHost, cfg.MQPort)
	if err != nil {
		log.Fatal(err)
	}
	authClient, err := client.NewAuthClient("localhost:" + cfg.GRPCPort)
	if err != nil {
		log.Fatal(err)
	}
	svc := service.NewOrderService(repo, authClient,rabbitmq)
	handler := handler.NewOrderHandler(svc)



	app := fiber.New()

	router.SetupRoutes(app, handler, cfg.JWTSecret)

	log.Println("Auth service running on port", cfg.AppPort)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
