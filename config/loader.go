package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get() *Config {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error when load env file %s", err)
	}

	return &Config{
		Redis: Redis{
			Address:  os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			Database: os.Getenv("REDIS_DB"),
		},
		Server: Server{
			Port: os.Getenv("SERVER_PORT"),
		},
	}
}
