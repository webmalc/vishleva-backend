package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/webmalc/vishleva-backend/dto"
)

// BookHandler is handler.
type BookHandler struct {
	booker Booker
}

// Post creates a new object.
func (h *BookHandler) Post(c *gin.Context) {
	request := dto.Book{
		Name:       "online order",
		ClientName: "online client",
	}
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err) // nolint // unnecessary

		return
	}
	if _, err := h.booker.Book(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err) // nolint // unnecessary

		return
	}
	c.JSON(http.StatusCreated, &request)
}

// NewBookHandler returns a new book handler  object.
func NewBookHandler(booker Booker) *BookHandler {
	return &BookHandler{booker: booker}
}
