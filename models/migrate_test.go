package models

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/models/mocks"
)

// Should migrate the DB.
func TestMigrate(t *testing.T) {
	am := &mocks.AutoMigrater{}
	conn := db.NewConnection()
	defer conn.Close()
	args := []interface{}{
		&User{},
		&Collection{},
		&Tag{},
		&Image{},
		&Tariff{},
		&Review{},
		&Client{},
		&Order{},
	}
	am.On("AutoMigrate", args...).Return(conn.DB).Once()
	am.On("Model", mock.Anything).Return(conn.DB).Times(4)
	Migrate(am)
	am.AssertExpectations(t)
}
