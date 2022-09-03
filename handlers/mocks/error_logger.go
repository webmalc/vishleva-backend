package mocks

import (
	"github.com/stretchr/testify/mock"
)

// ErrorLogger logs errors.
type ErrorLogger struct {
	mock.Mock
}

// Errorf is method mock.
func (m *ErrorLogger) Errorf(format string, args ...interface{}) {
	m.Called(format, args)
}
