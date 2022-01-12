package calendar

import (
	"github.com/spf13/viper"
)

// Config is the database configuration struct.
type Config struct {
	StartHour int
	EndHour   int
}

// setDefaults sets the default values
func setDefaults() {
	viper.SetDefault("calendar_start_hour", 7) // nolint // unnecessary: unparam
	viper.SetDefault("calendar_end_hour", 22)  // nolint // unnecessary: unparam
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		StartHour: viper.GetInt("calendar_start_hour"),
		EndHour:   viper.GetInt("calendar_end_hour"),
	}
	return config
}
