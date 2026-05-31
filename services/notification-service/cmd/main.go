package main

import (
	"notification-service/internal/config"
	"notification-service/internal/service"
)


func main() {
	cfg := config.LoadConfig()
	service.Consumer(cfg.MQUser,cfg.MQPassword,cfg.MQHost,cfg.MQPort)
}
