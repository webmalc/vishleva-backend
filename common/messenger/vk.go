package messenger

import (
	"strconv"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/pkg/errors"
)

// TODO: test it
// VkSender is a vk sender.
type VkSender struct {
	config      *Config
	vkAPISender VkAPISender
}

// Send sends message to the user.
func (s *VkSender) Send(id string, message MessageGetter, ch chan error) {

	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ch <- errors.Wrap(err, "vk")

		return
	}
	s.setVkAPISender()
	b := params.NewMessagesSendBuilder()
	b.Message(message.GetText())
	b.RandomID(0)
	b.PeerID(int(userID))

	_, err = s.vkAPISender.MessagesSend(b.Params)
	if err != nil {
		ch <- errors.Wrap(err, "vk")
	}
	ch <- nil
}

// setVkAPISender sets the VK API sender.
func (s *VkSender) setVkAPISender() {
	if s.vkAPISender == nil {
		s.vkAPISender = api.NewVK(s.config.VkToken)
	}
}

// NewVkSender creates a new vk sender.
func NewVkSender(config *Config) *VkSender {
	return &VkSender{config: config}
}
