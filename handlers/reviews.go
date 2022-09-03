package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReviewsHandler is handler
type ReviewsHandler struct {
	getter ReviewsGetter
}

// GetList returns the list handler function
func (h *ReviewsHandler) GetList(c *gin.Context) {
	reviews, _ := h.getter.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"entries": reviews,
	})
}

// NewReviewsHandler returns a new router object
func NewReviewsHandler(getter ReviewsGetter) *ReviewsHandler {
	return &ReviewsHandler{getter: getter}
}
