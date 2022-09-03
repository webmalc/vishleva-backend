package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// ListHandler mocks the object.
type ListHandler struct {
	mock.Mock
}

// GetList is method mock.
func (m *ListHandler) GetList(c *gin.Context) {
	m.Called(c)
}
