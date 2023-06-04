package main

import (
	"flag"

	"github.com/isutare412/goarch/http-base/pkg/config"
	"github.com/isutare412/goarch/http-base/pkg/log"
	"github.com/isutare412/goarch/http-base/pkg/metric"
	"github.com/isutare412/goarch/http-base/pkg/tracing"
	"github.com/isutare412/goarch/http-base/pkg/wire"
)

var configPath = flag.String("config", "config.yaml", "YAML config file path")

func init() {
	flag.Parse()
}

func main() {
	cfg, err := config.LoadValidated(*configPath)
	if err != nil {
		panic(err)
	}
	cfgHub := config.NewHub(cfg)

	log.Init(cfgHub.ToLogConfig())
	defer log.Sync()

	tracing.Init(cfgHub.ToTracingConfig())
	defer tracing.Shutdown()

	metric.Init()

	components, err := wire.NewComponents(cfgHub, cfg.Wire.ShutdownTimeout)
	if err != nil {
		log.WithOperation("wireComponents").Panicf("Failed to wire components: %v", err)
	}

	components.RunAndWait()
	components.GracefulShutdown()
}
