package config

import (
	"github.com/isutare412/goarch/api-server/pkg/controller/http"
	"github.com/isutare412/goarch/api-server/pkg/log"
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

func (h *Hub) ToHTTPServerConfig() http.Config {
	return http.Config{
		Host: h.cfg.HTTP.Host,
		Port: h.cfg.HTTP.Port,
	}
}
