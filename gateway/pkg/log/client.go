package log

import (
	"fmt"

	"github.com/isutare412/goarch/gateway/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	globalLogger        *zap.Logger        = zap.NewNop()
	globalSugaredLogger *zap.SugaredLogger = globalLogger.Sugar()
	globalAccessLogger  *zap.Logger        = zap.NewNop()
)

func Init(cfg config.LoggerConfig) error {
	var zCfg zap.Config
	if cfg.Format == config.LogFormatJson {
		zCfg = jsonLoggerConfig()
	} else {
		zCfg = textLoggerConfig()
	}
	zCfg.DisableStacktrace = !cfg.StackTrace
	zCfg.DisableCaller = !cfg.ReportCaller

	var logger *zap.Logger
	logger, err := zCfg.Build()
	if err != nil {
		return fmt.Errorf("building logger: %w", err)
	}

	var accessLogger *zap.Logger
	accessLogger, err = accessLoggerConfig(zCfg).Build()
	if err != nil {
		return fmt.Errorf("building access logger: %w", err)
	}

	globalLogger = logger
	globalSugaredLogger = logger.Sugar()
	globalAccessLogger = accessLogger
	return nil
}

func L() *zap.SugaredLogger {
	return globalSugaredLogger.With("type", "app")
}

func WithOperation(name string) *zap.SugaredLogger {
	return globalSugaredLogger.With(
		"type", "app",
		"operation", name,
	)
}

func A() *zap.Logger {
	return globalAccessLogger.With(zap.String("type", "access"))
}

func Sync() {
	globalLogger.Sync()
	globalSugaredLogger.Sync()
	globalAccessLogger.Sync()
}

func textLoggerConfig() zap.Config {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.CallerKey = "C"
	cfg.EncoderConfig.FunctionKey = "F"
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return cfg
}

func jsonLoggerConfig() zap.Config {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.CallerKey = "caller"
	cfg.EncoderConfig.FunctionKey = "func"
	return cfg
}

func accessLoggerConfig(cfg zap.Config) zap.Config {
	cfg.EncoderConfig.CallerKey = zapcore.OmitKey
	cfg.EncoderConfig.FunctionKey = zapcore.OmitKey
	return cfg
}
