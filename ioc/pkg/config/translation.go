package config

import (
	"github.com/isutare412/goarch/ioc/pkg/http"
	"github.com/isutare412/goarch/ioc/pkg/log"
)

func NewLogConfig(cfg Config) log.Config {
	return log.Config{
		Development: cfg.Log.Development,
		Format:      cfg.Log.Format,
		Level:       cfg.Log.Level,
		StackTrace:  cfg.Log.StackTrace,
		Caller:      cfg.Log.Caller,
	}
}

func NewHTTPConfig(cfg Config) http.Config {
	return http.Config{
		Port: cfg.HTTP.Port,
	}
}
