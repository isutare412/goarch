package log

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

type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
	LevelPanic Level = "panic"
	LevelFatal Level = "fatal"
)
