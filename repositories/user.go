package repositories

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/utils"
)

// UserRepository is the user repository.
type UserRepository struct {
	db *gorm.DB
}

// LoginAndReturnUser logins a user and return the user struct.
func (r *UserRepository) LoginAndReturnUser(
	email, password string,
) (*models.User, error) {
	var u models.User
	if r.db.Where(&models.User{Email: email}).First(&u).RecordNotFound() {
		return nil, errors.New("user is not found")
	}
	if !utils.CheckPassword(u.Password, password) {
		return nil, errors.New("password is incorrect")
	}
	now := time.Now()
	u.LastLogin = &now
	r.db.Save(&u)

	return &u, nil
}

// NewUserRepository return a new user repository struct.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
