package main

import (
	"flag"

	"github.com/isutare412/goarch/api-server/pkg/config"
	"github.com/isutare412/goarch/api-server/pkg/log"
	"github.com/isutare412/goarch/api-server/pkg/wire"
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

	components, err := wire.NewComponents(cfgHub)
	if err != nil {
		log.WithOperation("wireComponents").Fatalf("Failed to wire components: %v", err)
	}

	components.RunAndWait()
	components.GracefulShutdown()
}