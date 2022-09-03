package server

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/test"
)

// Should return the config object.
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, "9000", c.ServerPort)
	assert.Equal(t, "localhost", c.ServerHost)
	assert.False(t, c.IsReleaseMode)
	assert.Equal(t, 30*time.Second, c.ServerReadTimeout)
	assert.Equal(t, 30*time.Second, c.ServerWriteTimeout)
	assert.Equal(t, 5*time.Second, c.ServerShutdownTimeout)
	assert.Contains(t, c.ServerLogPath, "logs/server.test.log")
	assert.Contains(t, c.ServerAllowOrigins, "http://vishleva.com")
	assert.Contains(t, c.ServerAllowOrigins, "https://vishleva.com")
}

// Setups the tests.
func TestMain(m *testing.M) {
	if err := os.Chdir("../"); err != nil {
		panic(err)
	}
	test.Run(m)
}
