package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/models"
)

// CollectionsGetter mocks the object
type CollectionsGetter struct {
	mock.Mock
}

// GetAll is method mock
func (m *CollectionsGetter) GetAll() ([]models.Collection, []error) {
	arg := m.Called()
	return arg.Get(0).([]models.Collection), nil
}
