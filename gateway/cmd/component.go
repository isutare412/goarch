package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/isutare412/goarch/gateway/pkg/adapter/postgres"
	"github.com/isutare412/goarch/gateway/pkg/config"
	"github.com/isutare412/goarch/gateway/pkg/controller/http"
	"github.com/isutare412/goarch/gateway/pkg/core/service/account"
	"github.com/isutare412/goarch/gateway/pkg/core/service/meeting"
	"github.com/isutare412/goarch/gateway/pkg/log"
)

type components struct {
	postgresClient *postgres.Client
	httpServer     *http.Server
}

func (c *components) DependencyInjection(cfg config.Config) error {
	log.L().Info("Start dependency injection")

	postgresClient, err := postgres.NewClient(cfg.Postgres)
	if err != nil {
		return fmt.Errorf("creating PostgreSQL client: %w", err)
	}

	userRepository := postgres.NewUserRepository(postgresClient)
	adminRepository := postgres.NewAdminRepository(postgresClient)
	meetingRepository := postgres.NewMeetingRepository(postgresClient)

	accountService := account.ServiceBuilder().
		WithRepositorySession(postgresClient).
		WithUserRepository(userRepository).
		WithAdminRepository(adminRepository).
		Build()

	meetingService := meeting.ServiceBuilder().
		WithRepositorySession(postgresClient).
		WithUserRepository(userRepository).
		WithMeetingRepository(meetingRepository).
		Build()

	httpServer := http.ServerBuilder().
		WithHTTPServerConfig(cfg.Server.HTTP).
		WithAccountService(accountService).
		WithMeetingService(meetingService).
		Build()

	c.postgresClient = postgresClient
	c.httpServer = httpServer

	log.L().Info("Done dependency injection")
	return nil
}

func (c *components) Init(ctx context.Context) error {
	log.L().Info("Start components initialization")

	if err := c.postgresClient.MigrateSchemas(ctx); err != nil {
		return fmt.Errorf("migrating schemas: %w", err)
	}
	log.L().Info("Migrated database schemas")

	c.httpServer.Init()
	log.L().Info("Initialized HTTP server")

	log.L().Info("Done components initialization")
	return nil
}

func (c *components) RunAndWait() {
	httpServerFails := c.httpServer.Run()
	log.L().Infof("Running HTTP server from %q", c.httpServer.Addr())

	signals := make(chan os.Signal, 3)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case err := <-httpServerFails:
		log.L().Error(err)
	case sig := <-signals:
		log.L().Infof("Caught signal(%s)", sig.String())
	}
}

func (c *components) Shutdown(ctx context.Context) {
	log.L().Info("Start graceful shutdown")

	if err := c.httpServer.Shutdown(ctx); err != nil {
		log.L().Errorf("Failed to shutdown HTTP server: %v", err)
	}
	if err := c.postgresClient.Close(ctx); err != nil {
		log.L().Errorf("Failed to shutdown PostgreSQL client: %v", err)
	}

	log.L().Info("Done graceful shutdown")
}
