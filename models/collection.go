package models

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/qor/sorting"
)

// Collection is a model.
type Collection struct {
	gorm.Model
	sorting.Sorting
	Name        string `gorm:"size:255;not null;index;" valid:"required"`
	Summary     string `gorm:"type:text"`
	Description string `gorm:"type:text"`
	IsEnabled   bool   `gorm:"type:bool;default:false;index"`
	Tags        []*Tag `gorm:"many2many:collection_tags;"`
	ImageID     *uint
	Image       Image `gorm:"constraint:OnDelete:SET NULL;default:null"`
}

// MarshalJSON returns the JSON respresentation.
func (t *Collection) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
		Image       string `json:"image"`
	}{
		ID:          t.ID,
		Name:        t.Name,
		Summary:     t.Summary,
		Description: t.Description,
		Image:       t.Image.File.URL("middle"),
	})
}
