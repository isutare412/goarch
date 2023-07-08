package config

import "fmt"

type HTTPConfig struct {
	Port int `mapstructure:"port"`
}

func (c *HTTPConfig) Validate() error {
	if c.Port == 0 {
		return fmt.Errorf("port should not be empty")
	}
	return nil
}
