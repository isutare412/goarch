package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger = *zap.SugaredLogger

var (
	globalBaseLogger   Logger = zap.NewNop().Sugar()
	globalAppLogger    Logger = zap.NewNop().Sugar()
	globalAccessLogger Logger = zap.NewNop().Sugar()
)

func Init(cfg Config) {
	zcfg := baseZapConfig()
	zcfg.Encoding = cfg.Format.ZapEncoding()
	zcfg.Level = cfg.Level.ZapLevel()

	if cfg.Development {
		zcfg.Development = true
		zcfg.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder

		if cfg.Format == FormatText {
			zcfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		} else {
			zcfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		}
	} else {
		zcfg.Development = false
		zcfg.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
		zcfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	}

	zcfg.DisableStacktrace = !cfg.StackTrace
	zcfg.DisableCaller = !cfg.Caller

	logger, err := zcfg.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to build zap logger: %v", err))
	}

	setGlobalLogger(logger.Sugar())
}

func A() Logger {
	return globalAccessLogger
}

func L() Logger {
	return globalAppLogger
}

func WithOperation(op string) Logger {
	return globalAppLogger.With(
		zap.String("operation", op),
	)
}

func Sync() {
	globalBaseLogger.Sync()
}

func baseZapConfig() zap.Config {
	cfg := zap.NewProductionConfig()
	cfg.Sampling = nil
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return cfg
}

func setGlobalLogger(logger Logger) {
	globalBaseLogger = logger
	globalAppLogger = logger.With(zap.String("type", "app"))
	globalAccessLogger = logger.With(zap.String("type", "access")).WithOptions(zap.WithCaller(false))
}
