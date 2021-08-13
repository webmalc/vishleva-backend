package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ImagesHander is handler
type ImagesHander struct {
	getter ImagesGetter
}

// GetList returns the list handler function
func (h *ImagesHander) GetList(c *gin.Context) {
	collection := 0
	if v, err := strconv.Atoi(c.Query("collection")); err == nil && v > 0 {
		collection = v
	}
	images, _ := h.getter.GetAll(
		c.Query("tag"),
		uint(collection),
	)
	c.JSON(http.StatusOK, gin.H{
		"entries": images,
	})
}

// NewImagesHandler returns a new router object
func NewImagesHandler(getter ImagesGetter) *ImagesHander {
	return &ImagesHander{getter: getter}
}
