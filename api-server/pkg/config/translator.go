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
		Development: t.cfg.Logger.Development,
		Format:      t.cfg.Logger.Format,
		Level:       t.cfg.Logger.Level,
		StackTrace:  t.cfg.Logger.StackTrace,
		Caller:      t.cfg.Logger.Caller,
	}
}
