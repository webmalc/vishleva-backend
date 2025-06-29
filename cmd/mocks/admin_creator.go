package mocks

import (
	"github.com/stretchr/testify/mock"
)

// AdminCreator logs errors.
type AdminCreator struct {
	mock.Mock
}

// Error is method mock.
func (m *AdminCreator) CreateAdmin(email, password string) error {
	var _ca []interface{}
	_ca = append(_ca, email, password)
	m.Called(_ca...)

	return nil
}
