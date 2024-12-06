package mocks

import (
	"github.com/stretchr/testify/mock"
	"gopkg.in/gomail.v2"
)

// EmailDialer is mock object.
type EmailDialer struct {
	mock.Mock
}

// DialAndSend is method mock.
func (s *EmailDialer) DialAndSend(_ ...*gomail.Message) error {
	arg := s.Called()

	return arg.Get(0).(error)
}
