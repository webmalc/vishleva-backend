package models

import (
	"encoding/json"
	"errors"

	"github.com/jinzhu/gorm"
)

// Client is a model
type Client struct {
	gorm.Model
	Name    string  `gorm:"size:255;not null;index"`
	Comment string  `gorm:"type:text;index"`
	Email   *string `gorm:"size:255;index;unique;default:null" valid:"email"`
	Social  string  `gorm:"size:255;index"`
	Phone   string  `gorm:"size:255;index;" valid:"numeric"`
}

// Validate validates the client
func (t *Client) Validate(db *gorm.DB) {
	if t.Name == "" {
		_ = db.AddError(errors.New(
			"name is empty",
		))
	}
}

// MarshalJSON returns the JSON respresentation
func (t *Client) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Name)
}
