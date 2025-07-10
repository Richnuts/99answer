package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port     int    `envconfig:"PORT" default:"9002"`
	Database string `envconfig:"USER_SERVICE" default:"user.db"`
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	var conf Config
	err = envconfig.Process("", &conf)
	if err != nil {
		log.Printf("fail to proceed the config: %v", err)
	}

	return &conf
}
