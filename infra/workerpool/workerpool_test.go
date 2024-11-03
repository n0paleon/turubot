package workerpool

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
	"turubot/infra/config"
	logger2 "turubot/infra/logger"
)

func TestInitializePool(t *testing.T) {
	fp, err := filepath.Abs("../../config.yaml")
	assert.Nil(t, err, "cannot retrieve absolute path")

	cfg := config.LoadConfig(fp)
	assert.NotNil(t, cfg, "config is nil")

	logger := logger2.NewConsoleLogger(cfg)
	assert.NotNil(t, logger, "logger is nil")

	assert.NotNil(t, cfg.PoolSize)
	err = InitializePool(cfg.PoolSize, logger)
	assert.Nil(t, err, "cannot initialize pool")

	defer ClosePool()
}
