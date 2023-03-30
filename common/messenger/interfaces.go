package messenger

// ContactsGetter is the interface for contacts.
type ContactsGetter interface {
	GetEmail() string
	GetTelegram() string
	GetVK() string
}

// MessageGetter is the interface for messages.
type MessageGetter interface {
	GetText() string
	GetHTML() string
	GetSubject() string
}

// Sender is the interface for sending messages.
type Sender interface {
	Send(id string, message MessageGetter, ch chan error)
}

// ErrorLogger logs errors.
type ErrorLogger interface {
	Error(args ...interface{})
}

// ContactsConverter is the interface for converting contacts to map.
type ContactsConverter interface {
	Convert(contact ContactsGetter) map[string]string
}
