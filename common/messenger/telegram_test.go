package messenger

import (
	"fmt"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/common/messenger/mocks"
)

func TestTelegramSender_Send(t *testing.T) {
	telegram := NewTelegramSender(NewConfig())
	botSender := mocks.NewTelegramBotSender(t)
	botSender.On("Send", mock.Anything).Return(
		tgbotapi.Message{}, fmt.Errorf("test error"),
	).Once()
	telegram.botSender = botSender
	message := &mocks.MessageGetter{}
	message.On("GetText").Return("test text").Once()
	ch := make(chan error)
	go telegram.Send("111", message, ch)
	result := <-ch
	assert.Contains(t, result.Error(), "test error")

	message.AssertExpectations(t)
	botSender.AssertExpectations(t)
}

func TestTelegramSender_setBotSender(t *testing.T) {
	telegram := NewTelegramSender(NewConfig())
	err := telegram.setBotSender()
	assert.NotNil(t, err)

	botSender := mocks.NewTelegramBotSender(t)
	telegram.botSender = botSender
	err = telegram.setBotSender()
	assert.Nil(t, err)
	assert.Equal(t, botSender, telegram.botSender)
}

func TestNewTelegramSender(t *testing.T) {
	telegram := NewTelegramSender(NewConfig())
	assert.NotNil(t, telegram)
}
