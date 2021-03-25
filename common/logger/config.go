package logger

import (
	"github.com/spf13/viper"
)

// Config is the logger configuration struct.
type Config struct {
	FilePath string
	IsDebug  bool
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	config := &Config{
		FilePath: viper.GetString("base_dir") + viper.GetString("log_path"),
		IsDebug:  !viper.GetBool("is_prod"),
	}
	return config
}
