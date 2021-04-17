package server

import (
	"time"

	"github.com/spf13/viper"
)

// Config is the admin configuration struct.
type Config struct {
	ServerPort            string
	ServerHost            string
	ServerReadTimeout     time.Duration
	ServerWriteTimeout    time.Duration
	IsReleaseMode         bool
	ServerAllowOrigins    []string
	ServerLogPath         string
	ServerShutdownTimeout time.Duration
}

// setDefaults sets the default values
func setDefaults() {
	viper.SetDefault("server_port", "9000")
	viper.SetDefault("server_host", "localhost")
	viper.SetDefault("server_read_timeout", "30s")
	viper.SetDefault("server_write_timeout", "30s")
	viper.SetDefault("server_shutdown_timeout", "5s")
	viper.SetDefault("server_log_path", "logs/server.log")
	viper.SetDefault("server_allow_origins",
		"https://vishleva.com http://vishleva.com",
	)
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	logPath := viper.GetString("base_dir") + viper.GetString("server_log_path")
	config := &Config{
		ServerPort:            viper.GetString("server_port"),
		ServerHost:            viper.GetString("server_host"),
		IsReleaseMode:         viper.GetBool("is_prod"),
		ServerReadTimeout:     viper.GetDuration("server_read_timeout"),
		ServerWriteTimeout:    viper.GetDuration("server_write_timeout"),
		ServerAllowOrigins:    viper.GetStringSlice("server_allow_origins"),
		ServerShutdownTimeout: viper.GetDuration("server_shutdown_timeout"),
		ServerLogPath:         logPath,
	}
	return config
}
