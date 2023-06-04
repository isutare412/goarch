package wire

import (
	"github.com/isutare412/goarch/http-base/pkg/config"
	"github.com/isutare412/goarch/http-base/pkg/controller/http"
)

type controllers struct {
	httpServer   *http.Server
	metricServer *http.MetricServer
}

func (ctrl *controllers) wire(cfg *config.Hub) error {
	ctrl.wireHTTPServer(cfg)
	ctrl.wireMetricServer(cfg)
	return nil
}

func (ctrl *controllers) wireHTTPServer(cfg *config.Hub) {
	ctrl.httpServer = http.NewServer(cfg.ToHTTPServerConfig())
}

func (ctrl *controllers) wireMetricServer(cfg *config.Hub) {
	ctrl.metricServer = http.NewMetricServer(cfg.ToHTTPServerConfig())
}
