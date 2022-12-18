package config

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	AppName            string `envconfig:"APP_NAME" default:"dialogflow-api-and-opentelemetry"`
	Version            string `default:"1.0.0"`
	Debug              bool   `envconfig:"DEBUG" default:"false"`
	Port               int    `envconfig:"PORT" default:"8000"`
}

type Config interface {
	GetAppName() string
	GetVersion() string
	GetDebug() bool
	GetPort() int
}

func LoadConfig() (Config, error) {
	cfg := new(config)

	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *config) GetAppName() string {
	return c.AppName
}
func (c *config) GetVersion() string {
	return c.Version
}
func (c *config) GetDebug() bool {
	return c.Debug
}
func (c *config) GetPort() int {
	return c.Port
}