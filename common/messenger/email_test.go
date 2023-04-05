package messenger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/common/messenger/mocks"
)

func TestEmailSender_Send(t *testing.T) {
	dialer := &mocks.EmailDialer{}
	emailSender := NewEmailSender(NewConfig())
	emailSender.dialer = dialer

	dialer.On("DialAndSend", mock.Anything).Return(
		fmt.Errorf("test error"),
	).Once()
	message := &mocks.MessageGetter{}
	message.On("GetSubject").Return("test subject").Once()
	message.On("GetHTML").Return("test text").Once()

	ch := make(chan error)

	go emailSender.Send("test@test.com", message, ch)
	result := <-ch
	assert.Contains(t, result.Error(), "test error")

	message.AssertExpectations(t)
	dialer.AssertExpectations(t)
}

func TestNewEmailSender(t *testing.T) {
	emailSender := NewEmailSender(NewConfig())
	assert.NotNil(t, emailSender)
	assert.NotNil(t, emailSender.dialer)
}
