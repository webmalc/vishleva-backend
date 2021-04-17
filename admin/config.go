package admin

import (
	"github.com/spf13/viper"
)

// Config is the admin configuration struct.
type Config struct {
	AdminPath string
}

// setDefaults sets the default values
func setDefaults() {
	viper.SetDefault("admin_path", "admin")
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		AdminPath: viper.GetString("admin_path"),
	}
	return config
}
