package config

import (
	"github.com/isutare412/goarch/http-base/pkg/controller/http"
	"github.com/isutare412/goarch/http-base/pkg/log"
	"github.com/isutare412/goarch/http-base/pkg/tracing"
)

type Hub struct {
	cfg Config
}

func NewHub(cfg Config) *Hub {
	return &Hub{cfg: cfg}
}

func (h *Hub) ToLogConfig() log.Config {
	return log.Config{
		Development: h.cfg.Logger.Development,
		Format:      h.cfg.Logger.Format,
		Level:       h.cfg.Logger.Level,
		StackTrace:  h.cfg.Logger.StackTrace,
		Caller:      h.cfg.Logger.Caller,
	}
}

func (h *Hub) ToTracingConfig() tracing.Config {
	return tracing.Config{
		Enabled:                 h.cfg.Tracing.Enabled,
		ServiceName:             h.cfg.App,
		Environment:             h.cfg.Environment,
		JaegerCollectorEndpoint: h.cfg.Tracing.Jaeger.CollectorEndpoint,
	}
}

func (h *Hub) ToHTTPServerConfig() http.Config {
	return http.Config{
		Host: h.cfg.HTTP.Host,
		Port: h.cfg.HTTP.Port,
	}
}
