package wire

import (
	"github.com/isutare412/goarch/http-base/pkg/config"
	"github.com/isutare412/goarch/http-base/pkg/controller/http"
)

type controllers struct {
	httpServer *http.Server
}

func (ctrl *controllers) wire(cfg *config.Hub) error {
	ctrl.wireHTTPServer(cfg)
	return nil
}

func (ctrl *controllers) wireHTTPServer(cfg *config.Hub) {
	ctrl.httpServer = http.NewServer(cfg.ToHTTPServerConfig())
}
