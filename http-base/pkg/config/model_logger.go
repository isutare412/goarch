package config

import (
	"github.com/isutare412/goarch/http-base/pkg/log"
)

type LoggerConfig struct {
	Development bool       `mapstructure:"development"`
	Format      log.Format `mapstructure:"format"`
	Level       log.Level  `mapstructure:"level"`
	StackTrace  bool       `mapstructure:"stackTrace"`
	Caller      bool       `mapstructure:"caller"`
}

func (c *LoggerConfig) Validate() error {
	if err := c.Format.Validate(); err != nil {
		return err
	}
	if err := c.Level.Validate(); err != nil {
		return err
	}
	return nil
}
