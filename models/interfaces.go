package models

import "github.com/jinzhu/gorm"

// AutoMigrater auto migrate the DB
type AutoMigrater interface {
	AutoMigrate(values ...interface{}) *gorm.DB
}
