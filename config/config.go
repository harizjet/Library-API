package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		ServerConfig ServerConfig
		DBConfig     DBConfig
	}

	ServerConfig struct {
		Host string
		Port string
	}

	DBConfig struct {
		Host     string
		Port     string
		User     string
		Password string
		DBname   string
	}
)

func LoadConfig(file string) Config {

	err := godotenv.Load(file)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	return Config{
		ServerConfig: ServerConfig{
			Host: host,
			Port: port,
		},
		DBConfig: DBConfig{
			Host:     host,
			Port:     db_port,
			User:     db_user,
			Password: db_password,
			DBname:   db_name,
		},
	}
}
