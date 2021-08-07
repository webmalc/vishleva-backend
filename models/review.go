package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/sorting"
)

// Review is a model
type Review struct {
	gorm.Model
	sorting.Sorting
	Content   string `gorm:"type:text; index" valid:"required"`
	ClientID  *uint
	Client    Client `gorm:"constraint:OnDelete:SET NULL;default:null"`
	ImageID   *uint
	Image     Image `gorm:"constraint:OnDelete:SET NULL;default:null"`
	IsEnabled bool  `gorm:"type:bool;default:false;index"`
}
