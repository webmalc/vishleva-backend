package calendar

import "github.com/webmalc/vishleva-backend/models"

// AvailabilitySetter checks a slot availability
type AvailabilitySetter interface {
	Set(slot *models.CalendarSlot)
	Init()
}

// Initializer initializes the generator
type Initializer interface {
	Init(gen *Generator)
}

// OrdersGetter returns orders
type OrdersGetter interface {
	GetUpcoming() ([]models.Order, []error)
}
