package wire

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/isutare412/goarch/api-server/pkg/config"
	"github.com/isutare412/goarch/api-server/pkg/log"
)

type components struct {
	ctrl *controllers
}

func NewComponents(cfg *config.Hub) (*components, error) {
	var ctrl controllers
	if err := ctrl.wire(cfg); err != nil {
		return nil, fmt.Errorf("wiring controllers: %w", err)
	}

	return &components{
		ctrl: &ctrl,
	}, nil
}

func (c *components) RunAndWait() {
	log.WithOperation("lifecycle").Info("Run components")

	runnables := c.collectRunnables()

	runtimeErrs := make(chan error, len(runnables))
	for _, r := range runnables {
		errs := r.Run()
		go func(name string) {
			if err, ok := <-errs; ok {
				runtimeErrs <- fmt.Errorf("runtime error from %s: %w", name, err)
			}
		}(r.name)
	}

	signals := make(chan os.Signal, 3)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	log.WithOperation("lifecycle").Info("Wait for signal or runtime errors")
	select {
	case err := <-runtimeErrs:
		log.WithOperation("receiveRuntimeError").Error(err)
	case s := <-signals:
		log.WithOperation("receiveSignal").Infof("Caught signal %q", s)
	}
}

func (c *components) GracefulShutdown() {
	log.WithOperation("lifecycle").Info("Start graceful shutdown")
	defer log.WithOperation("lifecycle").Info("Done graceful shutdown")

	shutdownables := c.collectShutdownables()

	// TODO: Set context timeout.
	ctx := context.Background()

	for _, s := range shutdownables {
		if err := s.Shutdown(ctx); err != nil {
			log.WithOperation("shutdown").Errorf("Failed to shutdown %s: %v", s.name, err)
		}
	}
}

func (c *components) collectRunnables() []namedRunnable {
	var runs []namedRunnable
	runs = append(runs, asRunnable("httpServer", c.ctrl.httpServer))
	return runs
}

func (c *components) collectShutdownables() []namedShutdownable {
	var shuts []namedShutdownable
	shuts = append(shuts, asShutdownable("httpServer", c.ctrl.httpServer))
	return shuts
}
