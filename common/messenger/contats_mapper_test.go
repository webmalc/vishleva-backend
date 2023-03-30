package messenger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/messenger/mocks"
)

func TestContactsMapperConvert(t *testing.T) {
	converter := NewContactsConverter()
	contact := &mocks.ContactsGetter{}

	contact.On("GetEmail").Return("test@test.com").Once()
	contact.On("GetTelegram").Return("tg1").Once()
	contact.On("GetVK").Return("vk2").Once()
	result := converter.Convert(contact)
	assert.Equal(t, "test@test.com", result["email"])
	assert.Equal(t, "tg1", result["telegram"])
	assert.Equal(t, "vk2", result["vk"])
}

func TestNewContactsConverter(t *testing.T) {
	converter := NewContactsConverter()
	assert.NotNil(t, converter)
}
