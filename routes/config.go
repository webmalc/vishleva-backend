package routes

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config is the database configuration struct.
type Config struct {
	CacheTimeout time.Duration
}

// setDefaults sets the default values.
func setDefaults() {
	viper.SetDefault("cache_timeout_hours", fmt.Sprintf("%dh", 24*7)) // nolint // unnecessary: unparam
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		CacheTimeout: viper.GetDuration("cache_timeout_hours"),
	}

	return config
}
