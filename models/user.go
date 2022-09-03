package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/webmalc/vishleva-backend/services"
)

// User is the user struct.
type User struct {
	gorm.Model
	Email     string `gorm:"size:255;not null;index;unique" valid:"email,required"`
	Password  []byte
	LastLogin *time.Time
}

// SetPassword hashes and sets the provided password.
func (u *User) SetPassword(newPwd string) error {
	pwd, err := services.HashPassword([]byte(newPwd))
	if err != nil {
		return err
	}
	u.Password = pwd

	return nil
}

// DisplayName satisfies the interface for Qor Admin.
func (u *User) DisplayName() string {
	return u.Email
}
