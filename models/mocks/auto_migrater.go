package mocks

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
)

// AutoMigrater logs errors.
type AutoMigrater struct {
	mock.Mock
}

// AutoMigrate is a method mock
func (m *AutoMigrater) AutoMigrate(values ...interface{}) *gorm.DB {
	arg := m.Called(values...)
	return arg.Get(0).(*gorm.DB)
}
