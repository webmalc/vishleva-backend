package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Should return the config object.
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.NotNil(t, c)
}
