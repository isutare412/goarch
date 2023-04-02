package main

import (
	"flag"

	"github.com/isutare412/goarch/api-server/pkg/config"
	"github.com/isutare412/goarch/api-server/pkg/log"
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
	cfgTranslator := config.NewTranslator(cfg)

	log.Init(cfgTranslator.ToLogConfig())
	defer log.Sync()

	log.L().Info("this is the end")
}
