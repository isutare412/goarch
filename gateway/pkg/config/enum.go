package config

import "fmt"

type LogFormat string

const (
	LogFormatJson LogFormat = "json"
	LogFormatText LogFormat = "text"
)

func (lf LogFormat) Validate() error {
	switch lf {
	case LogFormatJson:
	case LogFormatText:
	default:
		return fmt.Errorf("unknown logFormat(%s)", lf)
	}
	return nil
}
