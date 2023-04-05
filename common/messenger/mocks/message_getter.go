package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MessageGetter is the interface for messages.
type MessageGetter struct {
	mock.Mock
}

// GetText is method mock.
func (s *MessageGetter) GetText() string {
	arg := s.Called()

	return arg.Get(0).(string)
}

// GetHTML is method mock.
func (s *MessageGetter) GetHTML() string {
	arg := s.Called()

	return arg.Get(0).(string)
}

// GetSubject is method mock.
func (s *MessageGetter) GetSubject() string {
	arg := s.Called()

	return arg.Get(0).(string)
}
