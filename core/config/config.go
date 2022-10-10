package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Conf *Config
)

type Config struct {
	AppVersion string
	JwtSecret  string
	AppPort    string
	Db         *ConfigDatabase
}

type ConfigDatabase struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func InitConfig() (*Config, error) {
	Conf := &Config{
		AppVersion: getEnvValue("APP_VERSION"),
		JwtSecret:  getEnvValue("JWT_SECRET"),
		AppPort:    getEnvValue("APP_PORT"),
		Db: &ConfigDatabase{
			Host:     getEnvValue("DB_HOST"),
			Port:     getEnvValue("DB_PORT"),
			Username: getEnvValue("DB_USERNAME"),
			Password: getEnvValue("DB_PASSWORD"),
			Database: getEnvValue("DB_DATABASE"),
		},
	}

	return Conf, nil
}

func getEnvValue(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
