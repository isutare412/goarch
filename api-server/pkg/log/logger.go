package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger = *zap.SugaredLogger

const (
	keyType      = "type"
	keyOperation = "operation"
)

const (
	typeApp    = "app"
	typeAccess = "access"
)

var (
	globalLogger Logger = zap.NewNop().Sugar()
)

func Init(cfg Config) {
	zcfg := baseZapConfig()

	switch cfg.Format {
	case FormatJSON:
		zcfg.Encoding = "json"
	case FormatText:
		zcfg.Encoding = "console"
	}

	switch cfg.Level {
	case LevelDebug:
		zcfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case LevelInfo:
		zcfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case LevelWarn:
		zcfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case LevelError:
		zcfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case LevelPanic:
		zcfg.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case LevelFatal:
		zcfg.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	}

	if cfg.Development {
		zcfg.Development = true
		zcfg.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	} else {
		zcfg.Development = false
		zcfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	}

	zcfg.DisableStacktrace = !cfg.StackTrace
	zcfg.DisableCaller = !cfg.Caller

	logger, err := zcfg.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to build zap logger: %v", err))
	}

	globalLogger = logger.Sugar()
}

func A() Logger {
	return globalLogger.With(zap.String(keyType, typeAccess))
}

func L() Logger {
	return globalLogger.With(zap.String(keyType, typeApp))
}

func WithOperation(op string) Logger {
	return globalLogger.With(
		zap.String(keyType, typeApp),
		zap.String(keyOperation, op),
	)
}

func Sync() {
	globalLogger.Sync()
}

func baseZapConfig() zap.Config {
	cfg := zap.NewProductionConfig()
	cfg.Sampling = nil
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return cfg
}
