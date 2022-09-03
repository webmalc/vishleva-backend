package handlers

import (
	"github.com/spf13/viper"
)

// Config is the configuration struct.
type Config struct {
	AdminPath string
	LoginPath string
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	config := &Config{
		AdminPath: viper.GetString("admin_path"),
		LoginPath: viper.GetString("login_path"),
	}

	return config
}
