package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// ListHander mocks the object
type ListHander struct {
	mock.Mock
}

// GetList is method mock
func (m *ListHander) GetList(c *gin.Context) {
	m.Called(c)
}
