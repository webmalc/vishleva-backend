package calendar

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/calendar/mocks"
)

func TestBaseInitializer_setDefaultBegin(t *testing.T) {
	og := &mocks.OrdersGetter{}
	ca := NewGenerator(og)
	i := BaseInitializer{}
	assert.True(t, ca.Begin.IsZero())
	i.setDefaultBegin(ca)
	assert.False(t, ca.Begin.IsZero())
}

func TestDayInitializer_Init(t *testing.T) {
	og := &mocks.OrdersGetter{}
	ca := NewGenerator(og)
	i := DayInitializer{}
	assert.True(t, ca.Begin.IsZero())
	assert.True(t, ca.End.IsZero())
	i.Init(ca)
	assert.False(t, ca.Begin.IsZero())
	assert.False(t, ca.End.IsZero())
	assert.Equal(t, ca.Begin, ca.End)
}

func TestWeekInitializer_Init(t *testing.T) {
	og := &mocks.OrdersGetter{}
	ca := NewGenerator(og)
	i := WeekInitializer{}
	assert.True(t, ca.Begin.IsZero())
	assert.True(t, ca.End.IsZero())
	i.Init(ca)
	assert.False(t, ca.Begin.IsZero())
	assert.False(t, ca.End.IsZero())
	assert.Equal(t, 6, int(ca.End.Sub(ca.Begin).Hours()/24))
}

func TestMonthInitializer_Init(t *testing.T) {
	og := &mocks.OrdersGetter{}
	ca := NewGenerator(og)
	i := MonthInitializer{}
	assert.True(t, ca.Begin.IsZero())
	assert.True(t, ca.End.IsZero())
	i.Init(ca)
	assert.False(t, ca.Begin.IsZero())
	assert.False(t, ca.End.IsZero())
	assert.GreaterOrEqual(t, int(ca.End.Sub(ca.Begin).Hours()/24), 28)
}

func TestYearInitializer_Init(t *testing.T) {
	og := &mocks.OrdersGetter{}
	ca := NewGenerator(og)
	i := YearInitializer{}
	assert.True(t, ca.Begin.IsZero())
	assert.True(t, ca.End.IsZero())
	ca.Begin = time.Date(2024, 1, 23, 14, 30, 0, 0, time.UTC)
	i.Init(ca)
	assert.False(t, ca.Begin.IsZero())
	assert.False(t, ca.End.IsZero())
	assert.Contains(t, []int{365, 366}, int(ca.End.Sub(ca.Begin).Hours()/24))
}
