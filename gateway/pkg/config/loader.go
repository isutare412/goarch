package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func Load(path string) (Config, error) {
	var cfg Config

	if err := readFile(path); err != nil {
		return cfg, err
	}
	readEnv()

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func LoadValidated(path string) (Config, error) {
	var cfg Config

	cfg, err := Load(path)
	if err != nil {
		return cfg, fmt.Errorf("loading config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return cfg, fmt.Errorf("validating config: %w", err)
	}
	return cfg, nil
}

func readFile(path string) error {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func readEnv() {
	// GATEWAY_FOO_BAR=baz -> cfg.Foo.Bar = "baz"
	viper.SetEnvPrefix("gateway")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
