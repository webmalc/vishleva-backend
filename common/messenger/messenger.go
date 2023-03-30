package messenger

import "github.com/pkg/errors"

// TODO: test it
// Messenger is the struct to send messages via different sources to contact.
type Messenger struct {
	config            *Config
	logger            ErrorLogger
	senders           map[string]Sender
	contactsConverter ContactsConverter
}

// Send the message to the contact.
func (s *Messenger) Send(
	contact ContactsGetter, message MessageGetter, sources ...string,
) {
	ch := make(chan error)
	sources = s.getAvailableSources(contact, sources)

	for _, id := range sources {
		go s.senders[id].Send(id, message, ch)
	}

	for range sources {
		err := <-ch
		if err != nil {
			s.logger.Error(errors.Wrap(err, "messenger"))
		}
	}
}

// getAvailableSources returns the available sources.
func (s *Messenger) getAvailableSources(
	contact ContactsGetter, sources []string,
) []string {
	if len(sources) == 0 {
		sources = s.config.Sources
	}
	contactsMap := s.contactsConverter.Convert(contact)
	availableSources := []string{}
	for _, source := range sources {
		id := contactsMap[source]
		if id == "" {
			continue
		}
		sender := s.senders[source]
		if sender == nil {
			continue
		}
		availableSources = append(availableSources, source)
	}

	return availableSources
}

// NewMessenger creates a new messenger.
func NewMessenger(logger ErrorLogger) *Messenger {
	config := NewConfig()
	senders := make(map[string]Sender)
	factory := NewSenderFactory()
	for _, id := range config.Sources {
		senders[id] = factory.New(id)
	}

	return &Messenger{
		config:            config,
		senders:           senders,
		logger:            logger,
		contactsConverter: NewContactsConverter(),
	}
}
