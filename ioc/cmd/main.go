package main

import (
	"flag"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/isutare412/goarch/ioc/pkg/config"
	"github.com/isutare412/goarch/ioc/pkg/core/service"
	"github.com/isutare412/goarch/ioc/pkg/http"
	"github.com/isutare412/goarch/ioc/pkg/log"
	"github.com/isutare412/goarch/ioc/pkg/memory"
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
	logger := log.NewLogger(config.NewLogConfig(cfg))

	fx.New(
		fx.Supply(cfg, logger),
		config.Module,
		memory.Module,
		service.Module,
		http.Module,
		fx.Invoke(
			func(*http.Server) {},
		),
		fx.RecoverFromPanics(),
		fx.StartTimeout(cfg.Lifecycle.StartTimeout),
		fx.StopTimeout(cfg.Lifecycle.StopTimeout),
		fx.WithLogger(func(log *log.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log.WithOperation("fx").Desugar()}
		}),
	).Run()
}
