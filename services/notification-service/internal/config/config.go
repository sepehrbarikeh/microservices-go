package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort string

	MQHost     string
	MQPort     string
	MQPassword string
	MQUser    string

	GRPCPort string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed to read config: ", err)
	}

	return &Config{
		AppPort: viper.GetString("app.port"),

		MQHost:     viper.GetString("rabbitmq.host"),
		MQPort:     viper.GetString("rabbitmq.port"),
		MQPassword: viper.GetString("rabbitmq.password"),
		MQUser:     viper.GetString("rabbitmq.user"),

		GRPCPort: viper.GetString("grpc.port"),
	}
}