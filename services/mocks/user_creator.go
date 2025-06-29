package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/models"
)

// UserCreator logs errors.
type UserCreator struct {
	mock.Mock
}

// UserCreator is method mock.
func (m *UserCreator) CreateUser(email, password string) (*models.User, error) {
	var _ca []any
	_ca = append(_ca, email, password)
	arg := m.Called(_ca...)

	return arg.Get(0).(*models.User), nil
}
