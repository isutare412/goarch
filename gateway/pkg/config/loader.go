package config

import (
	"strings"

	"github.com/spf13/viper"
)

func Load(path string) (*Config, error) {
	if err := readFile(path); err != nil {
		return nil, err
	}
	readEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
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
