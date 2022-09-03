package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TariffsHandler is handler.
type TariffsHandler struct {
	getter TariffsGetter
}

// GetList returns the list handler function.
func (h *TariffsHandler) GetList(c *gin.Context) {
	tariffs, _ := h.getter.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"entries": tariffs,
	})
}

// NewTariffsHandler returns a new router object.
func NewTariffsHandler(getter TariffsGetter) *TariffsHandler {
	return &TariffsHandler{getter: getter}
}
