package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	JWTSecret string

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

		DBHost:     viper.GetString("database.host"),
		DBPort:     viper.GetString("database.port"),
		DBUser:     viper.GetString("database.user"),
		DBPassword: viper.GetString("database.password"),
		DBName:     viper.GetString("database.name"),
		DBSSLMode:  viper.GetString("database.sslmode"),

		JWTSecret: viper.GetString("jwt.secret"),

		GRPCPort: viper.GetString("grpc.port"),
	}
}