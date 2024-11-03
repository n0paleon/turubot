package config

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestConfig(t *testing.T) {
	fp, err := filepath.Abs("../../config.yaml")
	assert.Nil(t, err, "cannot retrieve absolute path")

	cfg := LoadConfig(fp)
	assert.NotNil(t, cfg, "config is nil")

	assert.NotNil(t, cfg.App.Author)
	assert.Equal(t, "Muhammad Naufal Al Fattah", cfg.App.Author)
}
