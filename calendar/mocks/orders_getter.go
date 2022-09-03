package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/models"
)

// OrdersGetter mocks the object.
type OrdersGetter struct {
	mock.Mock
}

// GetAll is method mock.
func (m *OrdersGetter) GetUpcoming() ([]models.Order, []error) {
	arg := m.Called()

	return arg.Get(0).([]models.Order), nil
}
