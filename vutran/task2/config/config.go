package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

type Config struct {
	NodeEnv   string
	Database  DatabaseConfig
	JwtSecret string
}

var config *Config
var once sync.Once

func SetupConfig() *Config {
	once.Do(func() {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Something went wrong")
		}

		nodeEnv := os.Getenv("NODE_ENV")

		if nodeEnv == "" {
			log.Fatal("Missing node env")
		}

		jwtSecret := os.Getenv("JWT_SECRET")

		if jwtSecret == "" {
			log.Fatal("Missing jwt secret")
		}

		config = &Config{
			Database:  SetupDatabaseEnvironment(),
			NodeEnv:   nodeEnv,
			JwtSecret: jwtSecret,
		}

	})
	return config
}

func SetupDatabaseEnvironment() DatabaseConfig {
	host := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Port is invalid")
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return DatabaseConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DbName:   dbName,
	}
}
