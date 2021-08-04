package models

import (
	"github.com/jinzhu/gorm"
)

// Tag is a model
type Tag struct {
	gorm.Model
	Name        string        `gorm:"size:255;not null;index;unique" valid:"required"`
	Collections []*Collection `gorm:"many2many:collection_tags;"`
}
