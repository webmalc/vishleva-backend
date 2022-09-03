package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/test"
)

// Should hash the password.
func TestHashPassword(t *testing.T) {
	pass := []byte("password")
	hash, err := HashPassword(pass)
	assert.NotEqual(t, pass, hash)
	assert.Nil(t, err)

	_, err = HashPassword([]byte{})
	assert.Contains(t, err.Error(), "is empty")
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
