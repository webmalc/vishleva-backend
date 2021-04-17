package mocks

import (
	"github.com/stretchr/testify/mock"
)

// InfoLogger logs errors.
type InfoLogger struct {
	mock.Mock
}

// Infof is method mock
func (m *InfoLogger) Infof(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	m.Called(_ca...)
}
