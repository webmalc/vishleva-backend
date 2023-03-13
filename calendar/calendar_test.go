package calendar

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/repositories"
)

func TestGenerator_init(t *testing.T) {
	conn := db.NewConnection()
	repo := repositories.NewOrderRepository(conn.DB)
	c := NewGenerator(repo)
	c.init()
	assert.IsType(t, &WeekInitializer{}, c.Initializer)
}

func TestGenerator_setSlotAvailability(t *testing.T) {
	conn := db.NewConnection()
	repo := repositories.NewOrderRepository(conn.DB)
	now := time.Now()
	begin := time.Date(
		now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local,
	)
	c := NewGenerator(repo)

	// too early
	slot := &models.CalendarSlot{
		IsOpen: true, Begin: begin, End: begin.Add(time.Hour),
	}
	c.setSlotAvailability(slot)
	assert.False(t, slot.IsOpen)

	// midday
	slot.Begin = begin.Add(time.Hour * 12)
	slot.End = slot.Begin.Add(time.Hour)
	slot.IsOpen = true
	c.setSlotAvailability(slot)
	assert.True(t, slot.IsOpen)

	slot.IsOpen = false
	c.setSlotAvailability(slot)
	assert.False(t, slot.IsOpen)
}

func TestGenerator_generateSlots(t *testing.T) {
	conn := db.NewConnection()
	models.Migrate(conn)
	repo := repositories.NewOrderRepository(conn.DB)
	c := NewGenerator(repo)
	now := time.Now()
	begin := time.Date(
		now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, time.Local,
	)
	end := begin.Add(time.Hour)

	conn.Create(&models.Order{
		Name:   "test order",
		Begin:  &begin,
		End:    &end,
		Total:  decimal.NewFromInt(10),
		Paid:   decimal.NewFromInt(10),
		Status: "open",
		Source: "manual",
	})
	orders, _ := repo.GetUpcoming()

	for _, setter := range c.availabilitySetters {
		setter.Init()
	}
	slots := c.generateSlots(now)

	assert.Len(t, orders, 1)
	assert.Len(t, slots, 24)

	// working hours
	assert.False(t, slots[0].IsOpen)
	assert.False(t, slots[1].IsOpen)

	assert.True(t, slots[12].IsOpen)
	assert.True(t, slots[15].IsOpen)
	assert.False(t, slots[22].IsOpen)

	// order
	assert.False(t, slots[16].IsOpen)
	assert.True(t, slots[17].IsOpen)
}

func TestGenerator_Get(t *testing.T) {
	conn := db.NewConnection()
	models.Migrate(conn)
	repo := repositories.NewOrderRepository(conn.DB)
	generator := NewGenerator(repo)
	now := time.Now()
	result := generator.Get(now)
	weekLen := 7
	assert.Len(t, result, weekLen)

	// working hours
	assert.False(t, result[0].Slots[0].IsOpen)
	assert.True(t, result[2].Slots[10].IsOpen)
	assert.False(t, result[6].Slots[22].IsOpen)
}

func TestNewGenerator(t *testing.T) {
	conn := db.NewConnection()
	repo := repositories.NewOrderRepository(conn.DB)
	c := NewGenerator(repo)

	assert.Len(t, c.availabilitySetters, 2)
}
