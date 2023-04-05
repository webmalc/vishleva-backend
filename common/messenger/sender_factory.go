package messenger

// SenderFactory is a factory to create senders.
type SenderFactory struct{}

func (s *SenderFactory) New(id string, config *Config) Sender { // nolint // unnecessary: ireturn
	switch id {
	case "email":
		return NewEmailSender(config)
	case "telegram":
		return NewTelegramSender(config)
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
