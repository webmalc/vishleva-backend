package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// PostHandler mocks the object.
type PostHandler struct {
	mock.Mock
}

// Post is method mock.
func (m *PostHandler) Post(c *gin.Context) {
	m.Called(c)
}
