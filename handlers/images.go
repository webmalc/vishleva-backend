package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ImagesHandler is handler.
type ImagesHandler struct {
	getter ImagesGetter
}

// GetList returns the list handler function.
func (h *ImagesHandler) GetList(c *gin.Context) {
	collection := 0
	if v, err := strconv.Atoi(c.Query("collection")); err == nil && v > 0 {
		collection = v
	}
	images, _ := h.getter.GetAll(
		c.Query("tag"),
		uint(collection), // nolint // unnecessary: G115
	)
	c.JSON(http.StatusOK, gin.H{
		"entries": images,
	})
}

// NewImagesHandler returns a new router object.
func NewImagesHandler(getter ImagesGetter) *ImagesHandler {
	return &ImagesHandler{getter: getter}
}
