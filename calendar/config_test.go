package calendar

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/test"
)

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, c.StartHour, 7)
	assert.Equal(t, c.EndHour, 22)
}

// Setups the tests
func TestMain(m *testing.M) {
	test.Run(m)
}
