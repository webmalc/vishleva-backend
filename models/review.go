package models

import (
	"encoding/json"

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

// MarshalJSON returns the JSON respresentation
func (t *Review) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Content      string `json:"content"`
		Client       Client `json:"client"`
		Image        string `json:"image"`
		Duration     int    `json:"duration"`
		Photos       int    `json:"photos"`
		Retouch      int    `json:"retouch"`
		RetouchPrice string `json:"retouch_price"`
		IsPrimary    bool   `json:"is_primary"`
	}{
		Content: t.Content,
		Client:  t.Client,
		Image:   t.Image.File.URL("middle"),
	})
}
