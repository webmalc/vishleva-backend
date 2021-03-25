package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // according to the gorm docs
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // according to the gorm docs
	"github.com/pkg/errors"
	"github.com/qor/validations"
)

// Database is the database connection.
type Database struct {
	*gorm.DB
}

// NewConnection returns a new database connection.
func NewConnection() *Database {
	config := NewConfig()
	db, err := gorm.Open(config.DatabaseType, config.DatabaseURI)
	if err != nil {
		panic(errors.Wrap(err, "database"))
	}
	validations.RegisterCallbacks(db)
	return &Database{db}
}
