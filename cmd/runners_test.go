package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/cmd/mocks"
)

func TestAdminCreatorRunner_Run(t *testing.T) {
	ac := &mocks.AdminCreator{}
	runner := NewAdminCreatorRunner(ac)
	runner.passwordReader = func(int) ([]byte, error) {
		return []byte("password"), nil
	}
	ac.On("CreateAdmin", "email", "password").Return(nil).Once()
	runner.Run([]string{"email"})
	ac.AssertExpectations(t)
}

func TestNewAdminCreatorRunner(t *testing.T) {
	ac := &mocks.AdminCreator{}
	runner := NewAdminCreatorRunner(ac)
	assert.Equal(t, ac, runner.adminCreator)
}
