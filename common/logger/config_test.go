package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/test"
)

const filePathKey = "log_path"

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, true, c.IsDebug)
	assert.Contains(t, c.FilePath, "logs/app.test.log")
}

// Setups the tests
func TestMain(m *testing.M) {
	test.Run(m)
}
