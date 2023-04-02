package config

import "github.com/isutare412/goarch/api-server/pkg/log"

type Translator struct {
	cfg Config
}

func NewTranslator(cfg Config) *Translator {
	return &Translator{cfg: cfg}
}

func (t *Translator) ToLogConfig() log.Config {
	return log.Config{
		Development: t.cfg.LoggerConfig.Development,
		Format:      t.cfg.LoggerConfig.Format,
		Level:       t.cfg.LoggerConfig.Level,
		StackTrace:  t.cfg.LoggerConfig.StackTrace,
		Caller:      t.cfg.LoggerConfig.Caller,
	}
}
