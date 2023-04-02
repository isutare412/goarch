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
	zapCfg := baseZapConfig()

	switch cfg.Format {
	case FormatJSON:
		zapCfg.Encoding = "json"
	case FormatText:
		zapCfg.Encoding = "console"
	}

	switch cfg.Level {
	case LevelDebug:
		zapCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case LevelInfo:
		zapCfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case LevelWarn:
		zapCfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case LevelError:
		zapCfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case LevelPanic:
		zapCfg.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case LevelFatal:
		zapCfg.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	}

	if cfg.Development {
		zapCfg.Development = true
		zapCfg.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	} else {
		zapCfg.Development = false
		zapCfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	}

	zapCfg.DisableStacktrace = !cfg.StackTrace
	zapCfg.DisableCaller = !cfg.Caller

	logger, err := zapCfg.Build()
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
