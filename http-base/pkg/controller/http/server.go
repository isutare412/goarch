package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/isutare412/goarch/http-base/pkg/log"
)

type Server struct {
	srv  *http.Server
	addr string
}

func NewServer(cfg Config) *Server {
	devCtrl := devController{}

	r := chi.NewRouter()
	r.Use(middleware.RealIP, startTrace, requestLogger, recoverPanic)

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/dev", devCtrl.router())
	})

	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	return &Server{
		srv: &http.Server{
			Addr:    addr,
			Handler: r,
		},
		addr: addr,
	}
}

func (s *Server) Run() <-chan error {
	runtimeErrs := make(chan error, 1)
	go func() {
		log.WithOperation("runHTTPServer").Infof("Run HTTP server at '%s'", s.addr)

		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			runtimeErrs <- fmt.Errorf("HTTP server exited unexpectedly: %w", err)
		}
	}()
	return runtimeErrs
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdowning HTTP server: %w", err)
	}
	return nil
}
