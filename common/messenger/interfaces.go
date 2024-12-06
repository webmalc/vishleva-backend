package messenger

import (
	"gopkg.in/gomail.v2"

	"github.com/SevereCloud/vksdk/v2/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// ContactsGetter is the interface for contacts.
type ContactsGetter interface {
	GetEmail() string
	GetTelegram() string
	GetVK() string
}

// MessageGetter is the interface for messages.
type MessageGetter interface {
	GetText() string
	GetHTML() string
	GetSubject() string
}

// Sender is the interface for sending messages.
type Sender interface {
	Send(id string, message MessageGetter, ch chan error)
}

// Logger logs errors.
type Logger interface {
	Error(args ...interface{})
	Infof(format string, args ...interface{})
}

// ContactsConverter is the interface for converting contacts to map.
type ContactsConverter interface {
	Convert(contact ContactsGetter) map[string]string
}

// Dialer is the interface for dialing.
type EmailDialer interface {
	DialAndSend(m ...*gomail.Message) error
}

// TelegramBotSender is the interface for sending telegram messages.
type TelegramBotSender interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}

type VkAPISender interface {
	MessagesSend(b api.Params) (response int, err error)
}
