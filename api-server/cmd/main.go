package main

import "github.com/isutare412/goarch/api-server/pkg/log"

var logConfig = log.Config{
	Development: true,
	Format:      log.FormatText,
	Level:       log.LevelDebug,
	StackTrace:  false,
	Caller:      true,
}

func main() {
	log.Init(logConfig)
	defer log.Sync()
}
