package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/models"
)

// ImagesGetter mocks the object.
type ImagesGetter struct {
	mock.Mock
}

// GetAll is method mock.
func (m *ImagesGetter) GetAll(
	_ string, _ uint,
) ([]models.Image, []error) {
	arg := m.Called()

	return arg.Get(0).([]models.Image), nil
}
