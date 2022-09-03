package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CalendarHandler is handler.
type CalendarHandler struct {
	generator CalendarGenerator
}

// GetList returns the list handler function.
func (h *CalendarHandler) GetList(c *gin.Context) {
	t, _ := time.Parse(time.RFC3339, c.Query("date"))
	days := h.generator.Get(t)
	c.JSON(http.StatusOK, gin.H{
		"entries": days,
	})
}

// NewTagsHandler returns a new router object.
func NewCalendarHandler(getter CalendarGenerator) *CalendarHandler {
	return &CalendarHandler{generator: getter}
}
