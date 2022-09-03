package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// AuthHandler mocks the object
type AuthHandler struct {
	mock.Mock
}

// GetLogin is method mock
func (m *AuthHandler) GetLogin(c *gin.Context) {
	m.Called(c)
}

// GetLogout is method mock
func (m *AuthHandler) GetLogout(c *gin.Context) {
	m.Called(c)
}

// PostLogin is method mock
func (m *AuthHandler) PostLogin(c *gin.Context) {
	m.Called(c)
}
