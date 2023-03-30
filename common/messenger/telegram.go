package messenger

import (
	"fmt"
)

// TelegramSender is a telegram sender.
type TelegramSender struct{}

func (s *TelegramSender) Send(id string, message MessageGetter, ch chan error) {
	ch <- fmt.Errorf("telegram not implemented")
}

// NewTelegramSender creates a new telegram sender.
func NewTelegramSender() *TelegramSender {
	return &TelegramSender{}
}
