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
		PostgresDb: PostgresDb{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Dbname:   os.Getenv("POSTGRES_DB_NAME"),
		},
	}
}
