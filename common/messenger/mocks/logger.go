package mocks

import (
	"github.com/stretchr/testify/mock"
)

// Logger logs errors.
type Logger struct {
	mock.Mock
}

// Infof is method mock.
func (m *Logger) Infof(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	m.Called(_ca...)
}

// Error is method mock.
func (m *Logger) Error(args ...interface{}) {
	m.Called(args...)
}
