package config

type Config struct {
	Logger   LoggerConfig   `yaml:"logger"`
	Postgres PostgresConfig `yaml:"postgres"`
}

func (c Config) Validate() error {
	return c.Logger.Validate()
}

type LoggerConfig struct {
	Format     LogFormat `yaml:"format"`
	StackTrace bool      `yaml:"stackTrace"`
}

func (c LoggerConfig) Validate() error {
	return c.Format.Validate()
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
