package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/models"
)

// TariffsGetter mocks the object
type TariffsGetter struct {
	mock.Mock
}

// GetAll is method mock
func (m *TariffsGetter) GetAll() ([]models.Tariff, []error) {
	arg := m.Called()
	return arg.Get(0).([]models.Tariff), nil
}
