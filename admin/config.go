package admin

import (
	"github.com/spf13/viper"
)

// Config is the admin configuration struct.
type Config struct {
	AdminPath     string
	LoginPath     string
	LogoutPath    string
	SiteName      string
	OrderStatuses []string
}

// setDefaults sets the default values.
func setDefaults() {
	viper.SetDefault("admin_path", "admin")
	viper.SetDefault("site_name", "vishleva")
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		AdminPath:     viper.GetString("admin_path"),
		LoginPath:     viper.GetString("login_path"),
		LogoutPath:    viper.GetString("logout_path"),
		SiteName:      viper.GetString("site_name"),
		OrderStatuses: viper.GetStringSlice("order_statuses"),
	}

	return config
}
