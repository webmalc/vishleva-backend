package calendar

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/repositories"
)

func TestWorkingHoursAvailabilitySetter_Set(t *testing.T) {
	setter := NewWorkingHoursAvailabilitySetter()
	assert.Equal(
		t, setter.config.StartHour, viper.GetInt("calendar_start_hour"),
	)
	assert.Equal(
		t, setter.config.EndHour, viper.GetInt("calendar_end_hour"),
	)
}

func TestNewWorkingHoursAvailabilitySetter(t *testing.T) {
	setter := NewWorkingHoursAvailabilitySetter()
	now := time.Now()
	begin := time.Date(
		now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local,
	)

	// too early
	slot := &models.CalendarSlot{
		IsOpen: true, Begin: begin, End: begin.Add(time.Hour),
	}
	setter.Set(slot)
	assert.False(t, slot.IsOpen)

	// midday
	slot.Begin = begin.Add(time.Hour * 12)
	slot.End = slot.Begin.Add(time.Hour)
	slot.IsOpen = true
	setter.Set(slot)
	assert.True(t, slot.IsOpen)

	slot.IsOpen = false
	setter.Set(slot)
	assert.False(t, slot.IsOpen)
}

func TestOrdersAvailabilitySetter_Init(t *testing.T) {
	conn := db.NewConnection()
	models.Migrate(conn)
	repo := repositories.NewOrderRepository(conn.DB)
	setter := NewOrdersAvailabilitySetter(repo)

	// without order
	setter.Init()
	assert.Len(t, setter.closedHours, 0)

	// with order
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
	})

	setter.Init()
	assert.Len(t, setter.closedHours, 1)
	assert.True(t, setter.closedHours[begin.Format(setter.format)])
}

func TestOrdersAvailabilitySetter_Set(t *testing.T) {
	conn := db.NewConnection()
	models.Migrate(conn)
	repo := repositories.NewOrderRepository(conn.DB)
	setter := NewOrdersAvailabilitySetter(repo)

	now := time.Now()
	begin := time.Date(
		now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, time.Local,
	)
	end := begin.Add(time.Hour)
	slot := &models.CalendarSlot{
		IsOpen: true, Begin: begin, End: end,
	}
	setter.Init()
	setter.Set(slot)
	assert.True(t, slot.IsOpen)

	conn.Create(&models.Order{
		Name:   "test order",
		Begin:  &begin,
		End:    &end,
		Total:  decimal.NewFromInt(10),
		Paid:   decimal.NewFromInt(10),
		Status: "open",
	})

	setter.Init()
	setter.Set(slot)
	assert.False(t, slot.IsOpen)
}

func TestNewOrdersAvailabilitySetter(t *testing.T) {
	conn := db.NewConnection()
	repo := repositories.NewOrderRepository(conn.DB)
	setter := NewOrdersAvailabilitySetter(repo)
	assert.Equal(t, "2006-01-02 15:04", setter.format)
	assert.Equal(t, repo, setter.getter)
}
