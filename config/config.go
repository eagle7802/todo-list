package config

import (
	"fmt"
	logger "github.com/eagle7802/todo-list/pkg"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerConfig   ServerConfig
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     string `envconfig:"DB_PORT" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

type ServerConfig struct {
	Port string `envconfig:"SERVER_PORT" required:"true"`
}

func NewConfig(path string) (*Config, error) {
	if err := godotenv.Load(path); err != nil {
		logger.Errorf("godotenv.Load(): %v", err)
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	var config Config

	if err := envconfig.Process("", &config); err != nil {
		logger.Errorf("envconfig.Process(): %v", err)
		return nil, fmt.Errorf("error processing .env file: %v", err)
	}
	return &config, nil
}
