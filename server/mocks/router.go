package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// Router mocks the object.
type Router struct {
	mock.Mock
}

// BindRoutes is method mock.
func (m *Router) BindRoutes(e *gin.Engine) {
	m.Called(e)
}
