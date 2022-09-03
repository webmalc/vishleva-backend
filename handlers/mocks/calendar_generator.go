package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/models"
)

// CalendarGenerator mocks the object.
type CalendarGenerator struct {
	mock.Mock
}

// Get is a method mock.
func (m *CalendarGenerator) Get(begin time.Time) []*models.CalendarDay {
	arg := m.Called()
	return arg.Get(0).([]*models.CalendarDay)
}
