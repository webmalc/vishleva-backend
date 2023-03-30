package messenger

// SenderFactory is a factory to create senders.
type SenderFactory struct{}

func (s *SenderFactory) New(id string) Sender { // nolint // unnecessary: ireturn
	switch id {
	case "email":
		return NewEmailSender()
	case "telegram":
		return NewTelegramSender()
	case "vk":
		return NewVkSender()
	default:
		panic("SenderFactory: unknown sender")
	}
}

// NewSenderFactory creates a new sender factory.
func NewSenderFactory() *SenderFactory {
	return &SenderFactory{}
}
