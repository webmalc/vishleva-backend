package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/sorting"
)

// Review is a model
type Review struct {
	gorm.Model
	sorting.Sorting
	// client
	Content   string `gorm:"type:text; index" valid:"required"`
	ImageID   uint
	Image     Image `gorm:"constraint:OnDelete:SET NULL;"`
	IsEnabled bool  `gorm:"type:bool;default:false;index"`
}
