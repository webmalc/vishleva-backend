package messenger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/common/messenger/mocks"
)

// SenderMock logs errors.
type SenderMock struct {
	mock.Mock
}

// Send provides a mock function with given fields: id, message, ch.
func (m *SenderMock) Send(id string, message MessageGetter, ch chan error) {
	m.Called(id, message, ch)
	ch <- fmt.Errorf("test error")
}

func TestMessenger_Send(t *testing.T) {
	logger := &mocks.Logger{}
	logger.On("Infof", mock.Anything, "test@test.com", "email").Once()
	logger.On("Infof", mock.Anything, "tg1", "telegram").Once()
	logger.On("Infof", mock.Anything, "vk2", "vk").Once()
	logger.On("Error", mock.Anything).Times(3)
	messenger := NewMessenger(logger)

	contact := &mocks.ContactsGetter{}
	contact.On("GetEmail").Return("test@test.com").Once()
	contact.On("GetTelegram").Return("tg1").Once()
	contact.On("GetVK").Return("vk2").Once()

	message := &mocks.MessageGetter{}

	senderOne := &SenderMock{}
	senderTwo := &SenderMock{}
	senderThree := &SenderMock{}
	senderOne.On("Send", "test@test.com", message, mock.Anything).Return().Once()
	senderTwo.On("Send", "tg1", message, mock.Anything).Return().Once()
	senderThree.On("Send", "vk2", message, mock.Anything).Return().Once()

	messenger.senders = map[string]Sender{
		"email":    senderOne,
		"telegram": senderTwo,
		"vk":       senderThree,
	}

	messenger.Send(contact, message)
	senderOne.AssertExpectations(t)
	senderTwo.AssertExpectations(t)
	senderThree.AssertExpectations(t)
	contact.AssertExpectations(t)
	logger.AssertExpectations(t)
}

func TestMessenger_getAvailableSources(t *testing.T) {
	logger := &mocks.Logger{}
	messenger := NewMessenger(logger)
	result := messenger.getAvailableSources(
		map[string]string{"email": "email", "telegram": "tg1", "vk": "vk2"},
		[]string{"email", "telegram", "vk"},
	)
	assert.Equal(t, []string{"email", "telegram", "vk"}, result)

	result = messenger.getAvailableSources(
		map[string]string{"email": "", "telegram": "tg1"},
		[]string{"email", "telegram", "vk"},
	)
	assert.Equal(t, []string{"telegram"}, result)

	result = messenger.getAvailableSources(
		map[string]string{"email": "email", "telegram": "tg1", "vk": "vk2"},
		[]string{"vk"},
	)
	assert.Equal(t, []string{"vk"}, result)

	result = messenger.getAvailableSources(
		map[string]string{"email": "email", "telegram": "tg1", "vk": ""},
		[]string{"vk"},
	)
	assert.Empty(t, result)
}

func TestNewMessenger(t *testing.T) {
	logger := &mocks.Logger{}
	messenger := NewMessenger(logger)
	assert.NotNil(t, messenger)
	assert.NotNil(t, messenger.config)
	assert.NotNil(t, messenger.contactsConverter)
	assert.Len(t, messenger.senders, 3)
}
