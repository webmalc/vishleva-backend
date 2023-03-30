package messenger

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/test"
)

// Should return the config object.
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, viper.GetStringSlice("messenger_sources"), c.Sources)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
