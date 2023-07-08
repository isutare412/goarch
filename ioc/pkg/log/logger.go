package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	base   *zap.Logger
	app    *zap.SugaredLogger
	access *zap.Logger
}

func NewLogger(cfg Config) *Logger {
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

	l, err := zcfg.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to build zap logger: %v", err))
	}

	base := l
	return &Logger{
		base:   base,
		app:    base.Sugar().With(zap.String("type", "app")),
		access: base.With(zap.String("type", "access")).WithOptions(zap.WithCaller(false)),
	}
}

func (l *Logger) Access() *zap.Logger {
	return l.access
}

func (l *Logger) L() *zap.Logger {
	return l.app.Desugar()
}

func (l *Logger) S() *zap.SugaredLogger {
	return l.app
}

func (l *Logger) WithOperation(op string) *zap.SugaredLogger {
	return l.app.With(zap.String("operation", op))
}

func (l *Logger) Sync() {
	l.base.Sync()
}

func baseZapConfig() zap.Config {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout"}
	cfg.Sampling = nil
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return cfg
}
