package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/models"
)

// ReviewsGetter mocks the object.
type ReviewsGetter struct {
	mock.Mock
}

// GetAll is method mock.
func (m *ReviewsGetter) GetAll() ([]models.Review, []error) {
	arg := m.Called()

	return arg.Get(0).([]models.Review), nil
}
