package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/test"
)

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, c.ImageBigWidth, 144)
}

// Setups the tests
func TestMain(m *testing.M) {
	test.Run(m)
}
