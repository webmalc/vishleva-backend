package models

import (
	"testing"

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
	}
	am.On("AutoMigrate", args...).Return(conn.DB).Once()
	Migrate(am)
	am.AssertExpectations(t)
}
