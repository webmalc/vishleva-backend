package messenger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSenderFactoryNew(t *testing.T) {
	factory := NewSenderFactory()
	assert.IsType(t, &EmailSender{}, factory.New("email"))
	assert.IsType(t, &TelegramSender{}, factory.New("telegram"))
	assert.IsType(t, &VkSender{}, factory.New("vk"))
	assert.Panics(t, func() { factory.New("unknown") })
}

func TestNewSenderFactory(t *testing.T) {
	factory := NewSenderFactory()
	assert.NotNil(t, factory)
}
