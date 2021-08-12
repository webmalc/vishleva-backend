package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReviewsHander is handler
type ReviewsHander struct {
	getter ReviewsGetter
}

// GetList returns the list handler function
func (h *ReviewsHander) GetList(c *gin.Context) {
	tariffs, _ := h.getter.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"entries": tariffs,
	})
}

// NewReviewsHandler returns a new router object
func NewReviewsHandler(getter ReviewsGetter) *ReviewsHander {
	return &ReviewsHander{getter: getter}
}
