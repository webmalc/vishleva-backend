package handlers

import "github.com/webmalc/vishleva-backend/models"

// ErrorLogger logs errors.
type ErrorLogger interface {
	Errorf(format string, args ...interface{})
}

// UserLoginer logs errors.
type UserLoginer interface {
	LoginAndReturnUser(email, password string) (*models.User, error)
}
