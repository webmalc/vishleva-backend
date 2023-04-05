package messenger

import (
	"github.com/spf13/viper"
)

// Config is the configuration object.
type Config struct {
	Sources       []string
	EmailFrom     string
	EmailHost     string
	EmailPort     int
	EmailLogin    string
	EmailPassword string
	TelegramToken string
}

// setDefaults sets the default values.
func setDefaults() {
	viper.SetDefault("messenger_sources", []string{"email", "telegram", "vk"}) // nolint // unnecessary: unparam
	viper.SetDefault("messenger_email_port", 587)                              // nolint // unnecessary: unparam
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		Sources:       viper.GetStringSlice("messenger_sources"),
		EmailFrom:     viper.GetString("messenger_email_from"),
		EmailHost:     viper.GetString("messenger_email_host"),
		EmailPort:     viper.GetInt("messenger_email_port"),
		EmailLogin:    viper.GetString("messenger_email_login"),
		EmailPassword: viper.GetString("messenger_email_password"),
		TelegramToken: viper.GetString("messenger_telegram_token"),
	}

	return config
}
