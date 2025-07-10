package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port          int    `envconfig:"PORT" default:"9001"`
	ListingSvcURL string `envconfig:"LISTING_SERVICE" default:"http://localhost:6000"`
	UserSvcURL    string `envconfig:"USER_SERVICE" default:"http://localhost:9003"`
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
