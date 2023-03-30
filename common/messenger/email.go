package messenger

import (
	"fmt"
)

// EmailSender is a email sender.
type EmailSender struct{}

func (s *EmailSender) Send(id string, message MessageGetter, ch chan error) {
	ch <- fmt.Errorf("email not implemented")
}

// NewEmailSender creates a new email sender.
func NewEmailSender() *EmailSender {
	return &EmailSender{}
}
