package models

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Client is a model.
type Client struct {
	gorm.Model
	Name    string  `gorm:"size:255;not null;index"`
	Comment string  `gorm:"type:text;index"`
	Email   *string `gorm:"size:255;index;unique;default:null" valid:"email"`
	Social  string  `gorm:"size:255;index"`
	Phone   string  `gorm:"size:255;index;unique" valid:"numeric,length(11|11)"`
}

// Validate validates the client.
func (t *Client) Validate(db *gorm.DB) {
	c := NewConfig()
	phoneCode := fmt.Sprint(c.PhoneCode)
	if t.Name == "" {
		_ = db.AddError(errors.New(
			"name is empty",
		))
	}
	if t.Phone != "" && string(t.Phone[0]) != phoneCode {
		_ = db.AddError(errors.New(
			"phone must start with " + phoneCode,
		))
	}
}

// MarshalJSON returns the JSON representation.
func (t *Client) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Name)
}
