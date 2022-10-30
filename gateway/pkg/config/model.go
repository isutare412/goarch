package config

type Config struct {
	Logger   LoggerConfig   `yaml:"logger"`
	Server   ServerConfig   `yaml:"server"`
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

type ServerConfig struct {
	HTTP HTTPServerConfig `yaml:"http"`
}

type HTTPServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
