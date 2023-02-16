package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/webmalc/vishleva-backend/dto"
)

// TODO: test it
// BookHandler is handler.
type BookHandler struct{}

// Post creates a new object.
func (h *BookHandler) Post(c *gin.Context) {
	request := dto.Book{}
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// TODO:  run service
	fmt.Println(request)
	c.JSON(http.StatusAccepted, &request)
}

// NewBookHandler returns a new book handler  object.
func NewBookHandler() *BookHandler {
	return &BookHandler{}
}
