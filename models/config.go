package models

import (
	"github.com/spf13/viper"
)

// Config is the database configuration struct.
type Config struct {
	ImageSmallWidth   int
	ImageSmallHeight  int
	ImageMiddleWidth  int
	ImageMiddleHeight int
	ImageBigWidth     int
	ImageBigHeight    int
	OrderStatuses     []string
	PhoneCode         int
}

// setDefaults sets the default values.
func setDefaults() {
	viper.SetDefault("image_small_width", 60)    // nolint // unnecessary: unparam
	viper.SetDefault("image_small_height", 60)   // nolint // unnecessary: unparam
	viper.SetDefault("image_middle_width", 108)  // nolint // unnecessary: unparam
	viper.SetDefault("image_middle_height", 108) // nolint // unnecessary: unparam
	viper.SetDefault("image_big_width", 144)     // nolint // unnecessary: unparam
	viper.SetDefault("image_big_height", 144)    // nolint // unnecessary: unparam
	viper.SetDefault("phone_code", 7)            // nolint // unnecessary: unparam
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		ImageSmallWidth:   viper.GetInt("image_small_width"),
		ImageSmallHeight:  viper.GetInt("image_small_height"),
		ImageMiddleWidth:  viper.GetInt("image_middle_width"),
		ImageMiddleHeight: viper.GetInt("image_middle_height"),
		ImageBigWidth:     viper.GetInt("image_big_width"),
		ImageBigHeight:    viper.GetInt("image_big_height"),
		OrderStatuses:     viper.GetStringSlice("order_statuses"),
		PhoneCode:         viper.GetInt("phone_code"),
	}

	return config
}
