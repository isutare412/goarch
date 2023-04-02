package config

import "fmt"

type Config struct {
	Logger LoggerConfig `mapstructure:"logger"`
}

func (c *Config) Validate() error {
	if err := c.Logger.Validate(); err != nil {
		return fmt.Errorf("validating logger config: %w", err)
	}
	return nil
}
