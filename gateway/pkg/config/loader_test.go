package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const configPath = "../../configs/config.yaml"

func TestLoadFromFile(t *testing.T) {
	cfg, err := Load(configPath)
	require.NoError(t, err)

	require.NotNil(t, cfg)
	require.NotNil(t, cfg.Logger)
	assert.NotEmpty(t, cfg.Logger.Format)
}

func TestOverrideFromEnv(t *testing.T) {
	const (
		logFormat = "test"
	)

	t.Setenv("GATEWAY_LOGGER_FORMAT", logFormat)

	cfg, err := Load(configPath)
	require.NoError(t, err)

	require.NotNil(t, cfg)
	require.NotNil(t, cfg.Logger)
	assert.Equal(t, LogFormat(logFormat), cfg.Logger.Format)
}
