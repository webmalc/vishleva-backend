package routes

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/test"
)

// Should return the config object.
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, time.Hour*24*7, c.CacheTimeout)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
