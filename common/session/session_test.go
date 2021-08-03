package session

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSession(t *testing.T) {
	s := NewSession()
	assert.Equal(t, "vishleva_admin_session_test", s.Name)
	assert.Equal(t, "vishleva_user_id_test", s.Key)
	assert.NotNil(t, s.Store)
}
