package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CollectionHandler is handler
type CollectionHandler struct {
	getter CollectionsGetter
}

// GetList returns the list handler function
func (h *CollectionHandler) GetList(c *gin.Context) {
	collections, _ := h.getter.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"entries": collections,
	})
}

// NewCollectionHandler returns a new router object
func NewCollectionHandler(getter CollectionsGetter) *CollectionHandler {
	return &CollectionHandler{getter: getter}
}
