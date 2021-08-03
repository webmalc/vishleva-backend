package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/models"
)

// UserLogger mocks the object
type UserLogger struct {
	mock.Mock
}

// LoginAndReturnUser is method mock
func (m *UserLogger) LoginAndReturnUser(
	email, password string,
) (*models.User, error) {
	arg := m.Called(email, password)
	return arg.Get(0).(*models.User), nil
}
