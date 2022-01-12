package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CalendarHander is handler
type CalendarHander struct {
	generator CalendarGenerator
}

// GetList returns the list handler function
func (h *CalendarHander) GetList(c *gin.Context) {
	t, _ := time.Parse(time.RFC3339, c.Query("date"))
	days := h.generator.Get(t)
	c.JSON(http.StatusOK, gin.H{
		"entries": days,
	})
}

// NewTagsHandler returns a new router object
func NewCalendarHandler(getter CalendarGenerator) *CalendarHander {
	return &CalendarHander{generator: getter}
}
