package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/isutare412/goarch/http-base/pkg/log"
	"github.com/isutare412/goarch/http-base/pkg/metric"
)

type MetricServer struct {
	httpServer *http.Server
}

func NewMetricServer(cfg Config) *MetricServer {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(metric.Gatherer(), promhttp.HandlerOpts{}))

	return &MetricServer{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.MetricHost, cfg.MetricPort),
			Handler: mux,
		},
	}
}

func (srv *MetricServer) Run() <-chan error {
	fails := make(chan error, 1)
	go func() {
		defer close(fails)

		log.WithOperation("runMetricServer").Infof("Runn metric server on %s", srv.httpServer.Addr)

		err := srv.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fails <- fmt.Errorf("metrics server listen: %w", err)
			return
		}
	}()
	return fails
}

func (srv *MetricServer) Shutdown(ctx context.Context) error {
	return srv.httpServer.Shutdown(ctx)
}
