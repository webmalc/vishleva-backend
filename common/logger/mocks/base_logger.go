package mocks

import (
	"github.com/stretchr/testify/mock"
)

// BaseLogger is mock object
type BaseLogger struct {
	mock.Mock
}

// Debug is method mock
func (m *BaseLogger) Debug(args ...interface{}) {
	m.Called(args...)
}

// Debugf is method mock
func (m *BaseLogger) Debugf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	m.Called(_ca...)
}

// Info is method mock
func (m *BaseLogger) Info(args ...interface{}) {
	m.Called(args...)
}

// Infof is method mock
func (m *BaseLogger) Infof(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	m.Called(_ca...)
}

// Error is method mock
func (m *BaseLogger) Error(args ...interface{}) {
	m.Called(args...)
}

// Errorf is method mock
func (m *BaseLogger) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	m.Called(_ca...)
}

// Fatal is method mock
func (m *BaseLogger) Fatal(args ...interface{}) {
	m.Called(args...)
}

// Fatalf is method mock
func (m *BaseLogger) Fatalf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	m.Called(_ca...)
}
