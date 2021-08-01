package repositories

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/services"
)

// TODO: test it
type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) LoginAndReturnUser(
	email, password string,
) (*models.User, error) {
	var u models.User
	if r.db.Where(&models.User{Email: email}).First(&u).RecordNotFound() {
		return nil, errors.New("user is not found")
	}
	if !services.CheckPassword(u.Password, password) {
		return nil, errors.New("password is incorrect")
	}
	now := time.Now()
	u.LastLogin = &now
	r.db.Save(&u)

	return &u, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
