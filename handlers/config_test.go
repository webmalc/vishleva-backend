package handlers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/test"
)

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, "admin", c.AdminPath)
	assert.Equal(t, "/auth/login", c.LoginPath)
}

// Setups the tests
func TestMain(m *testing.M) {
	if err := os.Chdir("../"); err != nil {
		panic(err)
	}
	test.Run(m)
}
