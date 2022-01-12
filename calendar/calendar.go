package calendar

import (
	"time"

	"github.com/webmalc/vishleva-backend/models"
)

// TODO: test it
// Generator is the calendar generator
type Generator struct {
	Begin               time.Time
	End                 time.Time
	Initializer         Initializer
	availabilitySetters []AvailabilitySetter
}

// init setups the params
func (c *Generator) init() {
	if c.Initializer == nil {
		c.Initializer = &WeekInitializer{}
	}
	c.Initializer.Init(c)
}

// setSlotAvailability sets a slot availability
func (c *Generator) setSlotAvailability(slot *models.CalendarSlot) {
	for _, setter := range c.availabilitySetters {
		if !slot.IsOpen {
			return
		}
		setter.Set(slot)
	}
}

// generateSlots generate the slots
func (c *Generator) generateSlots(d time.Time) []*models.CalendarSlot {
	m := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local)
	n := m.AddDate(0, 0, 1)
	result := []*models.CalendarSlot{}
	for d := m; !d.Equal(n); d = d.Add(time.Hour) {
		slot := models.CalendarSlot{
			Begin:  d,
			End:    d.Add(time.Hour),
			IsOpen: true,
		}
		c.setSlotAvailability(&slot)
		result = append(result, &slot)
	}
	return result
}

// Get returns the result
func (c *Generator) Get(begin time.Time) []*models.CalendarDay {
	c.Begin = begin
	c.init()
	for _, setter := range c.availabilitySetters {
		setter.Init()
	}
	result := []*models.CalendarDay{}
	for d := c.Begin; !d.After(c.End); d = d.AddDate(0, 0, 1) {
		day := models.CalendarDay{
			Day:       d,
			IsWeekend: d.Weekday() == time.Sunday || d.Weekday() == time.Saturday,
			Slots:     c.generateSlots(d),
		}
		result = append(result, &day)
	}
	return result
}

// NewGenerator returns a new server object
func NewGenerator(getter OrdersGetter) *Generator {
	calendar := Generator{availabilitySetters: []AvailabilitySetter{
		NewWorkingHoursAvailalabilitySetter(),
		NewOrdersAvailalabilitySetter(getter),
	}}
	return &calendar
}
