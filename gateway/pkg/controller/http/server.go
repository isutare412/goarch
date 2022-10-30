package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isutare412/goarch/gateway/pkg/config"
	"github.com/isutare412/goarch/gateway/pkg/core/port"
	"github.com/isutare412/goarch/gateway/pkg/log"
)

type Server struct {
	cfg config.HTTPServerConfig

	accSvc port.AccountService
	mtgSvc port.MeetingService

	userHdr  *userHandler
	adminHdr *adminHandler

	engine *gin.Engine
	srv    *http.Server
}

func (s *Server) Init() {
	s.initEngine()
	s.initHandlers()
	s.initRoutes()

	s.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port),
		Handler: s.engine,
	}
}

func (s *Server) Run() <-chan error {
	fails := make(chan error, 1)
	go func() {
		err := s.srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fails <- fmt.Errorf("listening from http server: %w", err)
			return
		}
	}()
	return fails
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) Addr() string { return s.srv.Addr }

func (s *Server) initEngine() {
	gin.SetMode(gin.ReleaseMode)
	s.engine = gin.New()
}

func (s *Server) initHandlers() {
	s.userHdr = &userHandler{
		accSvc: s.accSvc,
	}

	s.adminHdr = &adminHandler{
		accSvc: s.accSvc,
	}
}

func (s *Server) initRoutes() {
	// TODO: Add middlewares
	api := s.engine.Group("/api/v1")
	s.userHdr.registerRoutes(api)
	s.adminHdr.registerRoutes(api)

	s.engine.GET("/dev", func(c *gin.Context) {
		log.L().With("headers", c.Request.Header).Info("Dev API called")
	})
}
