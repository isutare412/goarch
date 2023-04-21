package log

import (
	"fmt"

	"go.uber.org/zap"
)

type Config struct {
	Development bool
	Format      Format
	Level       Level
	StackTrace  bool
	Caller      bool
}

type Format string

const (
	FormatJSON Format = "json"
	FormatText Format = "text"
)

func (f Format) Validate() error {
	switch f {
	case FormatJSON, FormatText:
		return nil
	default:
		return fmt.Errorf("invalid log format '%s'", f)
	}
}

func (f Format) ZapEncoding() string {
	var zf string
	switch f {
	case FormatJSON:
		zf = "json"
	case FormatText:
		fallthrough
	default:
		zf = "console"
	}
	return zf
}

type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
	LevelPanic Level = "panic"
	LevelFatal Level = "fatal"
)

func (l Level) Validate() error {
	switch l {
	case LevelDebug, LevelInfo, LevelWarn, LevelError, LevelPanic, LevelFatal:
		return nil
	default:
		return fmt.Errorf("invalid log level '%s'", l)
	}
}

func (l Level) ZapLevel() zap.AtomicLevel {
	var zlevel zap.AtomicLevel
	switch l {
	case LevelDebug:
		zlevel = zap.NewAtomicLevelAt(zap.DebugLevel)
	case LevelInfo:
		zlevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	case LevelWarn:
		zlevel = zap.NewAtomicLevelAt(zap.WarnLevel)
	case LevelError:
		zlevel = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case LevelPanic:
		zlevel = zap.NewAtomicLevelAt(zap.PanicLevel)
	case LevelFatal:
		zlevel = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		zlevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	return zlevel
}
