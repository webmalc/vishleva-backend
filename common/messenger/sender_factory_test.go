package messenger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSenderFactoryNew(t *testing.T) {
	factory := NewSenderFactory()
	config := NewConfig()
	assert.IsType(t, &EmailSender{}, factory.New("email", config))
	assert.IsType(t, &TelegramSender{}, factory.New("telegram", config))
	assert.IsType(t, &VkSender{}, factory.New("vk", config))
	assert.Panics(t, func() { factory.New("unknown", config) })
}

func TestNewSenderFactory(t *testing.T) {
	factory := NewSenderFactory()
	assert.NotNil(t, factory)
}
