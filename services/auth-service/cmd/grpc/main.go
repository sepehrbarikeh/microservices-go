package main

import (
	"auth-service/internal/config"
	"auth-service/internal/db"
	server "auth-service/internal/grpc"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"fmt"
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

	repo := repository.NewUserRepository(database)
	jwtSvc := service.NewJWTService(cfg.JWTSecret)
	userSvc := service.NewUserService(repo, jwtSvc)

	server.Serve(userSvc, cfg.GRPCPort)
}