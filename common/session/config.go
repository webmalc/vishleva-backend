package session

import (
	"github.com/spf13/viper"
)

// Config is the admin configuration struct.
type Config struct {
	SessionKey    string
	SessionName   string
	SessionSecret string
}

// setDefaults sets the default values
func setDefaults() {
	viper.SetDefault("session_name", "vishleva_admin_session")
	viper.SetDefault("session_key", "vishleva_user_id")
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		SessionKey:    viper.GetString("session_key"),
		SessionName:   viper.GetString("session_name"),
		SessionSecret: viper.GetString("secret"),
	}
	return config
}
