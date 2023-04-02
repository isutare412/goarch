package log

import "fmt"

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
