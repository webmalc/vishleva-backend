package messenger

// TODO: test it
// ContactsMapper maps contacts to the sources.
type ContactsMapper struct{}

// Convert converts contacts to map.
func (c *ContactsMapper) Convert(contact ContactsGetter) map[string]string {
	contacts := make(map[string]string)
	contacts["email"] = contact.GetEmail()
	contacts["telegram"] = contact.GetTelegram()
	contacts["vk"] = contact.GetVK()

	return contacts
}

// NewContactsConverter creates a new contacts mapper.
func NewContactsConverter() *ContactsMapper {
	return &ContactsMapper{}
}
