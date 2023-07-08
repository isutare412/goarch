package config

import "github.com/isutare412/goarch/ioc/pkg/log"

type LogConfig struct {
	Development bool       `mapstructure:"development"`
	Format      log.Format `mapstructure:"format"`
	Level       log.Level  `mapstructure:"level"`
	StackTrace  bool       `mapstructure:"stackTrace"`
	Caller      bool       `mapstructure:"caller"`
}

func (c *LogConfig) Validate() error {
	if err := c.Format.Validate(); err != nil {
		return err
	}
	if err := c.Level.Validate(); err != nil {
		return err
	}
	return nil
}
