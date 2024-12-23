package messenger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/common/messenger/mocks"
)

func TestVkSender_Send(t *testing.T) {
	vk := NewVkSender(NewConfig())
	apiSender := mocks.NewVkAPISender(t)
	apiSender.On("MessagesSend", mock.Anything).Return(
		0, fmt.Errorf("test error"),
	).Once()
	vk.vkAPISender = apiSender
	message := &mocks.MessageGetter{}
	message.On("GetText").Return("test text").Once()
	ch := make(chan error)
	go vk.Send("111", message, ch)
	result := <-ch
	assert.Contains(t, result.Error(), "test error")

	message.AssertExpectations(t)
	apiSender.AssertExpectations(t)
}

func TestVkSender_setVkAPISender(t *testing.T) {
	vk := NewVkSender(NewConfig())
	vk.setVkAPISender()
	assert.NotNil(t, vk.vkAPISender)

	apiSender := mocks.NewVkAPISender(t)
	vk.vkAPISender = apiSender
	vk.setVkAPISender()
	assert.NotNil(t, vk.vkAPISender)
	assert.Equal(t, apiSender, vk.vkAPISender)
}

func TestNewVkSender(t *testing.T) {
	vk := NewVkSender(NewConfig())
	assert.NotNil(t, vk)
}
