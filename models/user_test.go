package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Must set the user password
func TestUser_SetPassword(t *testing.T) {
	u := User{}
	err := u.SetPassword("password")
	assert.Nil(t, err)
	assert.NotNil(t, u.Password)
	assert.NotEqual(t, []byte("password"), u.Password)

	err = u.SetPassword("")
	assert.NotNil(t, err)
}

// Must return the user name
func TestUser_DisplayName(t *testing.T) {
	u := User{}
	u.Email = "test@example.com"
	assert.Equal(t, u.Email, u.DisplayName())
}
