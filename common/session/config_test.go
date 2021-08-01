package session

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/test"
)

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, "vishleva_admin_session", c.SessionName)
	assert.Equal(t, "vishleva_user_id", c.SessionKey)
	assert.Equal(t, "secret_password_test", c.SessionSecret)
}

// Setups the tests
func TestMain(m *testing.M) {
	test.Run(m)
}
