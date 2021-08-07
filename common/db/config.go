package db

import (
	"github.com/spf13/viper"
)

// Config is the database configuration struct.
type Config struct {
	DatabaseURI  string
	DatabaseType string
}

// setDefaults sets the default values
func setDefaults() {
	d := "host=localhost port=5432 "
	d += "user=postgres dbname=vishleva_backend password=postgres"
	d += " TimeZone=UTC"
	viper.SetDefault("database_uri", d)
	viper.SetDefault("database_type", "postgres")
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		DatabaseURI:  viper.GetString("database_uri"),
		DatabaseType: viper.GetString("database_type"),
	}
	return config
}
