package messenger

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
)

// TelegramSender is a telegram sender.
type TelegramSender struct {
	config    *Config
	botSender TelegramBotSender
}

// Send sends a message to the channel.
func (s *TelegramSender) Send(id string, message MessageGetter, ch chan error) {
	err := s.setBotSender()
	if err != nil {
		ch <- err

		return
	}

	chatID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ch <- errors.Wrap(err, "telegram")

		return
	}
	m := tgbotapi.NewMessage(chatID, message.GetText())
	_, err = s.botSender.Send(m)
	if err != nil {
		ch <- errors.Wrap(err, "telegram")

		return
	}

	ch <- nil
}

// setBotSender sets the bot sender.
func (s *TelegramSender) setBotSender() error {
	if s.botSender != nil {
		return nil
	}
	bot, err := tgbotapi.NewBotAPI(s.config.TelegramToken)
	if err != nil {
		return errors.Wrap(err, "telegram")
	}
	s.botSender = bot

	return nil
}

// NewTelegramSender creates a new telegram sender.
func NewTelegramSender(config *Config) *TelegramSender {
	return &TelegramSender{config: config}
}
