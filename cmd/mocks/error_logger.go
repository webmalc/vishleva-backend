package mocks

import (
	"github.com/stretchr/testify/mock"
)

// ErrorLogger logs errors.
type ErrorLogger struct {
	mock.Mock
}

// Error is method mock
func (m *ErrorLogger) Error(args ...interface{}) {
	m.Called(args...)
}
