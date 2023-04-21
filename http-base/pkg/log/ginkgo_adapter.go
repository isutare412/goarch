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

func AdaptGinkgo() {
	encCfg := zap.NewProductionConfig().EncoderConfig
	encCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encCfg.EncodeDuration = zapcore.StringDurationEncoder
	encCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	zenc := zapcore.NewConsoleEncoder(encCfg)
	zcore := zapcore.NewCore(zenc, ginkgoSyncer{ginkgo.GinkgoWriter}, zapcore.DebugLevel)
	setGlobalLogger(zap.New(zcore, zap.AddCaller()).Sugar())
}
