package config

import "fmt"

type TracingConfig struct {
	Enabled bool         `mapstructure:"enabled"`
	Jaeger  JaegerConfig `mapstructure:"jaeger"`
}

func (c *TracingConfig) Validate() error {
	if !c.Enabled {
		return nil
	}

	if err := c.Jaeger.Validate(); err != nil {
		return fmt.Errorf("validating jaeger config: %w", err)
	}
	return nil
}

type JaegerConfig struct {
	CollectorEndpoint string `mapstructure:"collectorEndpoint"`
}

func (c *JaegerConfig) Validate() error {
	if c.CollectorEndpoint == "" {
		return fmt.Errorf("collectorEndpoint should not be empty")
	}
	return nil
}
