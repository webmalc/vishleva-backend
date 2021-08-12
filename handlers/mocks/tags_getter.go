package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/models"
)

// TagsGetter mocks the object
type TagsGetter struct {
	mock.Mock
}

// GetAll is method mock
func (m *TagsGetter) GetAll() ([]models.Tag, []error) {
	arg := m.Called()
	return arg.Get(0).([]models.Tag), nil
}
