package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TagsHander is handler
type TagsHander struct {
	getter TagsGetter
}

// GetList returns the list handler function
func (h *TagsHander) GetList(c *gin.Context) {
	tags, _ := h.getter.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"entries": tags,
	})
}

// NewTagsHandler returns a new router object
func NewTagsHandler(getter TagsGetter) *TagsHander {
	return &TagsHander{getter: getter}
}
