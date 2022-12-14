package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/isutare412/goarch/gateway/pkg/config"
	"github.com/isutare412/goarch/gateway/pkg/log"
)

var cfgPath = flag.String("config", "configs/config.yaml", "path to yaml config file")

// @Title Goarch Gateway
// @Version 1.0
// @Description Gateway of Goarch system.
func main() {
	flag.Parse()

	var cfg config.Config
	cfg, err := config.LoadValidated(*cfgPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load valid config: %v", err))
	}

	if err := log.Init(cfg.Logger); err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}
	defer log.Sync()

	var cmps components
	if err := cmps.DependencyInjection(cfg); err != nil {
		log.L().Fatalf("Failed during denpendency injection: %v", err)
	}

	initCtx, cancel := context.WithTimeout(context.Background(), cfg.Main.InitTimeout)
	defer cancel()

	if err := cmps.Init(initCtx); err != nil {
		log.L().Fatalf("Initializing components: %v", err)
	}

	cmps.RunAndWait()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Main.ShutdownTimeout)
	defer cancel()

	cmps.Shutdown(shutdownCtx)
}
