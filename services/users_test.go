package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/services/mocks"
)

func TestUsersService_CreateAdmin(t *testing.T) {
	log := &mocks.InfoLogger{}
	creator := &mocks.UserCreator{}
	service := NewUsersService(log, creator)
	user := &models.User{Email: "test@test.com"}
	creator.On(
		"CreateUser", mock.Anything, mock.Anything).Return(user, nil).Once()
	log.On("Infof", mock.Anything, user.Email).Return(nil).Once()
	err := service.CreateAdmin("email", "password")

	creator.AssertExpectations(t)
	log.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestNewUsersService(t *testing.T) {
	log := &mocks.InfoLogger{}
	creator := &mocks.UserCreator{}
	service := NewUsersService(log, creator)
	assert.Equal(t, log, service.logger)
	assert.Equal(t, creator, service.userCreator)
}
