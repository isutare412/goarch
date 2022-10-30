package config

import (
	"fmt"
	"time"
)

type Config struct {
	Main     MainConfig     `yaml:"main"`
	Logger   LoggerConfig   `yaml:"logger"`
	Server   ServerConfig   `yaml:"server"`
	Postgres PostgresConfig `yaml:"postgres"`
}

func (c Config) Validate() error {
	if err := c.Main.Validate(); err != nil {
		return err
	}
	if err := c.Logger.Validate(); err != nil {
		return err
	}
	return nil
}

type MainConfig struct {
	InitTimeout     time.Duration `yaml:"initTimeout"`
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout"`
}

func (c MainConfig) Validate() error {
	if c.InitTimeout == 0 {
		return fmt.Errorf("initTimeout should not be zero")
	}
	if c.ShutdownTimeout == 0 {
		return fmt.Errorf("shutdownTimeout should not be zero")
	}
	return nil
}

type LoggerConfig struct {
	Format       LogFormat `yaml:"format"`
	ReportCaller bool      `yaml:"reportCaller"`
	StackTrace   bool      `yaml:"stackTrace"`
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
