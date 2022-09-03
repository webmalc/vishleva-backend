package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/handlers/mocks"
)

func TestCalendarHandler_GetList(t *testing.T) {
	checkResponse(t, "/api/calendar", 7)
}

func TestNewCalendarHandler(t *testing.T) {
	cg := &mocks.CalendarGenerator{}
	handler := NewCalendarHandler(cg)
	assert.Equal(t, cg, handler.generator)
}
