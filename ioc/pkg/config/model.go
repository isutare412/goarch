package config

import (
	"fmt"
)

type Config struct {
	App         string          `mapstructure:"app"`
	Environment string          `mapstructure:"environment"`
	Version     string          `mapstructure:"version"`
	Lifecycle   LifecycleConfig `mapstructure:"lifecycle"`
	Log         LogConfig       `mapstructure:"log"`
	HTTP        HTTPConfig      `mapstructure:"http"`
}

func (c *Config) Validate() error {
	if c.App == "" {
		return fmt.Errorf("app should not be empty")
	}
	if c.Environment == "" {
		return fmt.Errorf("environment should not be empty")
	}
	if c.Version == "" {
		return fmt.Errorf("version should not be empty")
	}
	if err := c.Lifecycle.Validate(); err != nil {
		return fmt.Errorf("validating lifecycle config: %w", err)
	}
	if err := c.Log.Validate(); err != nil {
		return fmt.Errorf("validating log config: %w", err)
	}
	if err := c.HTTP.Validate(); err != nil {
		return fmt.Errorf("validating http config: %w", err)
	}
	return nil
}
