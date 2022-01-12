package calendar

import (
	"time"

	"github.com/webmalc/vishleva-backend/models"
)

// WorkingHoursAvailabilitySetter is the availability setter
type WorkingHoursAvailabilitySetter struct {
	config *Config
}

// Init initializes the setter
func (s *WorkingHoursAvailabilitySetter) Init() {}

// Set sets a slot availability
func (s *WorkingHoursAvailabilitySetter) Set(slot *models.CalendarSlot) {
	if slot.Begin.Hour() < s.config.StartHour {
		slot.IsOpen = false
		return
	}
	if s.config.EndHour > 0 && slot.Begin.Hour() >= s.config.EndHour {
		slot.IsOpen = false
	}
}

// NewWorkingHoursAvailalabilitySetter return a new object
func NewWorkingHoursAvailalabilitySetter() *WorkingHoursAvailabilitySetter {
	return &WorkingHoursAvailabilitySetter{config: NewConfig()}
}

// OrdersAvailabilitySetter is the availability setter
type OrdersAvailabilitySetter struct {
	getter      OrdersGetter
	closedHours map[string]bool
	format      string
}

func (s *OrdersAvailabilitySetter) Init() {
	orders, _ := s.getter.GetUpcoming()
	for i := range orders {
		o := &orders[i]
		begin := time.Date(
			o.Begin.Year(), o.Begin.Month(), o.Begin.Day(),
			o.Begin.Hour(), 0, 0, 0, time.Local,
		)
		end := *o.End
		if end.Minute() > 0 {
			end = time.Date(
				end.Year(), end.Month(), end.Day(),
				end.Hour()+1, 0, 0, 0, time.Local,
			)
		}
		for d := begin; !d.Equal(end); d = d.Add(time.Hour) {
			s.closedHours[d.Format(s.format)] = true
		}
	}
}

// Set sets a slot availability
func (s *OrdersAvailabilitySetter) Set(slot *models.CalendarSlot) {
	if s.closedHours[slot.Begin.Format(s.format)] {
		slot.IsOpen = false
	}
}

// NewWorkingHoursAvailalabilitySetter return a new object
func NewOrdersAvailalabilitySetter(
	getter OrdersGetter,
) *OrdersAvailabilitySetter {
	o := &OrdersAvailabilitySetter{
		getter:      getter,
		format:      "2006-01-02 15:04",
		closedHours: make(map[string]bool),
	}
	return o
}
