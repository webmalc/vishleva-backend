package admin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/test"
)

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, "admin", c.AdminPath)
}

// Setups the tests
func TestMain(m *testing.M) {
	test.Run(m)
}
