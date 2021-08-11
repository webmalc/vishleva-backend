package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TariffsHander is handler
type TariffsHander struct {
	getter TariffsGetter
}

// GetList returns the list handler function
func (h *TariffsHander) GetList(c *gin.Context) {
	tariffs, _ := h.getter.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"entries": tariffs,
	})
}

// NewTariffsHandler returns a new router object
func NewTariffsHandler(getter TariffsGetter) *TariffsHander {
	return &TariffsHander{getter: getter}
}
