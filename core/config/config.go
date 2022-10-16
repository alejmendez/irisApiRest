package config

import (
	"os"
)

var (
	Conf *Config
)

type Config struct {
	AppVersion   string
	JwtSecret    string
	JwtExpiresIn string
	AppPort      string
	Db           *ConfigDatabase
}

type ConfigDatabase struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func InitConfig() (*Config, error) {
	Config := &Config{
		AppVersion:   getEnvValue("APP_VERSION"),
		JwtSecret:    getEnvValue("JWT_SECRET"),
		JwtExpiresIn: getEnvValue("JWT_EXPIRES_IN"),
		AppPort:      getEnvValue("APP_PORT"),
		Db: &ConfigDatabase{
			Host:     getEnvValue("DB_HOST"),
			Port:     getEnvValue("DB_PORT"),
			Username: getEnvValue("DB_USERNAME"),
			Password: getEnvValue("DB_PASSWORD"),
			Database: getEnvValue("DB_DATABASE"),
		},
	}

	Conf = Config
	return Config, nil
}

func getEnvValue(key string) string {
	return os.Getenv(key)
}
