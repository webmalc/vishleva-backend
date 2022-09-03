package models

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

// Tag is a model.
type Tag struct {
	gorm.Model
	Name        string        `gorm:"size:255;not null;index;unique" valid:"required"`
	Collections []*Collection `gorm:"many2many:collection_tags;"`
}

// MarshalJSON returns the JSON respresentation.
func (t *Tag) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Name)
}
