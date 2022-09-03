package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TagsHandler is handler.
type TagsHandler struct {
	getter TagsGetter
}

// GetList returns the list handler function.
func (h *TagsHandler) GetList(c *gin.Context) {
	tags, _ := h.getter.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"entries": tags,
	})
}

// NewTagsHandler returns a new router object.
func NewTagsHandler(getter TagsGetter) *TagsHandler {
	return &TagsHandler{getter: getter}
}
