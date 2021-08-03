package mocks

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

// Admin mocks the object
type Admin struct {
	mock.Mock
}

// Mount is method mock
func (m *Admin) Mount() *http.ServeMux {
	arg := m.Called()
	return arg.Get(0).(*http.ServeMux)
}

// GetBasePath is method mock
func (m *Admin) GetBasePath() string {
	arg := m.Called()
	return arg.Get(0).(string)
}
