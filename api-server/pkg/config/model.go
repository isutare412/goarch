package config

import "fmt"

type Config struct {
	LoggerConfig LoggerConfig `mapstructure:"logger"`
}

func (c *Config) Validate() error {
	if err := c.LoggerConfig.Validate(); err != nil {
		return fmt.Errorf("validating logger config: %w", err)
	}
	return nil
}
