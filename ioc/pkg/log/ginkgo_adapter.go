package log

import (
	"github.com/onsi/ginkgo/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ginkgoSyncer struct {
	ginkgo.GinkgoWriterInterface
}

func (gs ginkgoSyncer) Sync() error { return nil }

func NewGinkgoLogger() *Logger {
	encCfg := zap.NewProductionConfig().EncoderConfig
	encCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encCfg.EncodeDuration = zapcore.StringDurationEncoder
	encCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	zenc := zapcore.NewConsoleEncoder(encCfg)
	zcore := zapcore.NewCore(zenc, ginkgoSyncer{ginkgo.GinkgoWriter}, zapcore.DebugLevel)

	l := zap.New(zcore, zap.AddCaller())
	base := l
	return &Logger{
		base:   base,
		app:    base.Sugar().With(zap.String("type", "app")),
		access: base.With(zap.String("type", "access")).WithOptions(zap.WithCaller(false)),
	}
}
