package config

import "fmt"

type MetricConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func (c *MetricConfig) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host should not be empty")
	}
	if c.Port == 0 {
		return fmt.Errorf("port should not be empty")
	}
	return nil
}
