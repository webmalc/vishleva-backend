package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/sorting"
)

// Collection is a model
type Collection struct {
	gorm.Model
	sorting.Sorting
	Name        string `gorm:"size:255;not null;index;" valid:"required"`
	Summary     string `gorm:"type:text"`
	Description string `gorm:"type:text"`
	IsEnabled   bool   `gorm:"type:bool;default:false;index"`
	Tags        []*Tag `gorm:"many2many:collection_tags;"`
}
