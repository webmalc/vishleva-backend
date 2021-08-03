package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

// AuthHander mocks the object
type AuthHander struct {
	mock.Mock
}

// GetLogin is method mock
func (m *AuthHander) GetLogin(c *gin.Context) {
	m.Called(c)
}

// GetLogout is method mock
func (m *AuthHander) GetLogout(c *gin.Context) {
	m.Called(c)
}

// PostLogin is method mock
func (m *AuthHander) PostLogin(c *gin.Context) {
	m.Called(c)
}
