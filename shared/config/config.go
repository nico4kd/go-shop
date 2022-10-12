package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Configurator interface {
	SetEnvironmentVariables() error
	GetDatabaseConfig() DatabaseConfig
	GetPort() string
}

type DatabaseConfig struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"54321"`
	Username string `env:"DB_USERNAME" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:"postgres"`
	Name     string `env:"DB_NAME" envDefault:"postgres"`
}

type CustomConfigurator struct {
	Database    DatabaseConfig
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
	Port        string `env:"PORT" envDefault:"3000"`
}

func NewCustomConfigurator() *CustomConfigurator {
	err := godotenv.Load()
	if err != nil {
		return &CustomConfigurator{}
	}
	config := CustomConfigurator{}
	if err := env.Parse(&config); err != nil {
		return &CustomConfigurator{}
	}
	return &config
}

var _ Configurator = (*CustomConfigurator)(nil)

func (c *CustomConfigurator) GetDatabaseConfig() DatabaseConfig {
	return c.Database
}

func (c *CustomConfigurator) SetEnvironmentVariables() error {
	//TODO implement me
	panic("implement me")
}

func (c *CustomConfigurator) GetPort() string {
	return c.Port
}
