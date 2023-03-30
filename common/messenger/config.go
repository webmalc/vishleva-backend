package messenger

import (
	"github.com/spf13/viper"
)

// TODO: test it
// Config is the configuration object.
type Config struct {
	Sources []string
}

// setDefaults sets the default values.
func setDefaults() {
	viper.SetDefault("messenger_sources", []string{"email", "telegram", "vk"}) // nolint // unnecessary: unparam
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		Sources: viper.GetStringSlice("messenger_sources"),
	}

	return config
}
