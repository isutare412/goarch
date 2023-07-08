package config

import (
	"fmt"
	"time"
)

type LifecycleConfig struct {
	StartTimeout time.Duration `mapstructure:"startTimeout"`
	StopTimeout  time.Duration `mapstructure:"stopTimeout"`
}

func (c *LifecycleConfig) Validate() error {
	if c.StartTimeout == 0 {
		return fmt.Errorf("startTimeout should not be empty")
	}
	if c.StopTimeout == 0 {
		return fmt.Errorf("stopTimeout should not be empty")
	}
	return nil
}
