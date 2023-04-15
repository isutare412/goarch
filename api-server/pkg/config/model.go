package config

import "fmt"

type Config struct {
	Logger LoggerConfig `mapstructure:"logger"`
	HTTP   HTTPConfig   `mapstructure:"http"`
}

func (c *Config) Validate() error {
	if err := c.Logger.Validate(); err != nil {
		return fmt.Errorf("validating logger config: %w", err)
	}
	if err := c.HTTP.Validate(); err != nil {
		return fmt.Errorf("validating http config: %w", err)
	}
	return nil
}
