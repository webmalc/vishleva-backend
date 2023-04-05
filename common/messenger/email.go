package messenger

import (
	"github.com/pkg/errors"
	"gopkg.in/gomail.v2"
)

// EmailSender is a email sender.
type EmailSender struct {
	config *Config
	dialer EmailDialer
}

// Send sends a message.
func (s *EmailSender) Send(id string, message MessageGetter, ch chan error) {
	email := gomail.NewMessage()
	email.SetHeader("From", s.config.EmailFrom)
	email.SetHeader("To", id)
	email.SetHeader("Subject", message.GetSubject())
	email.SetBody("text/html", message.GetHTML())

	if err := s.dialer.DialAndSend(email); err != nil {
		ch <- errors.Wrap(err, "email")

		return
	}
	ch <- nil
}

// NewEmailSender creates a new email sender.
func NewEmailSender(config *Config) *EmailSender {
	return &EmailSender{
		config: config,
		dialer: gomail.NewDialer(
			config.EmailHost,
			config.EmailPort,
			config.EmailLogin,
			config.EmailPassword,
		),
	}
}
