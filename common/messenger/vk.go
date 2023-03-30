package messenger

import (
	"fmt"
)

// VkSender is a vk sender.
type VkSender struct{}

func (s *VkSender) Send(id string, message MessageGetter, ch chan error) {
	ch <- fmt.Errorf("vk not implemented")
}

// NewVkSender creates a new vk sender.
func NewVkSender() *VkSender {
	return &VkSender{}
}
