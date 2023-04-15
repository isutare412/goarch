package config

import (
	"fmt"
	"time"
)

type Config struct {
	Wire   WireConfig   `mapstructure:"wire"`
	Logger LoggerConfig `mapstructure:"logger"`
	HTTP   HTTPConfig   `mapstructure:"http"`
}

func (c *Config) Validate() error {
	if err := c.Wire.Validate(); err != nil {
		return fmt.Errorf("validating wire config: %w", err)
	}
	if err := c.Logger.Validate(); err != nil {
		return fmt.Errorf("validating logger config: %w", err)
	}
	if err := c.HTTP.Validate(); err != nil {
		return fmt.Errorf("validating http config: %w", err)
	}
	return nil
}

type WireConfig struct {
	ShutdownTimeout time.Duration `mapstructure:"shutdownTimeout"`
}

func (c *WireConfig) Validate() error {
	if c.ShutdownTimeout == 0 {
		return fmt.Errorf("shutdown timeout should not be empty")
	}
	return nil
}
