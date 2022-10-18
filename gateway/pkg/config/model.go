package config

type Config struct {
	Logger *LoggerConfig `yaml:"logger"`
}

func (c *Config) Validate() error {
	return c.Logger.Validate()
}

type LoggerConfig struct {
	Format     LogFormat `yaml:"format"`
	StackTrace bool      `yaml:"stackTrace"`
}

func (c *LoggerConfig) Validate() error {
	return c.Format.Validate()
}
