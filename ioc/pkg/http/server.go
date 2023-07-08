package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"

	"github.com/isutare412/goarch/ioc/pkg/log"
)

type Server struct {
	server *http.Server
	log    *log.Logger
}

func NewServer(
	lc fx.Lifecycle,
	shut fx.Shutdowner,
	cfg Config,
	log *log.Logger,
	customerHandler *customerHandler,
) *Server {
	r := chi.NewRouter()
	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/customers", customerHandler.router())
	})

	srv := &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Port),
			Handler: r,
		},
		log: log,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error { return srv.Start(ctx, shut) },
		OnStop:  srv.Stop,
	})
	return srv
}

func (s *Server) Start(ctx context.Context, shut fx.Shutdowner) error {
	go func() {
		s.log.S().Infof("Starting HTTP server at %s", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.WithOperation("httpListenAndServe").Errorf("Failed to listen: %v", err)
			shut.Shutdown()
		}
	}()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdowning HTTP server: %w", err)
	}
	return nil
}
