package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CollectionHander is handler
type CollectionHander struct {
	getter CollectionsGetter
}

// GetList returns the list handler function
func (h *CollectionHander) GetList(c *gin.Context) {
	tariffs, _ := h.getter.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"entries": tariffs,
	})
}

// NewCollectionHandler returns a new router object
func NewCollectionHandler(getter CollectionsGetter) *CollectionHander {
	return &CollectionHander{getter: getter}
}
