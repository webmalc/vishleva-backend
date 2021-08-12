package handlers

import "github.com/webmalc/vishleva-backend/models"

// ErrorLogger logs errors.
type ErrorLogger interface {
	Errorf(format string, args ...interface{})
}

// UserLoginer logs users
type UserLoginer interface {
	LoginAndReturnUser(email, password string) (*models.User, error)
}

// TariffsGetter gets entries
type TariffsGetter interface {
	GetAll() ([]models.Tariff, []error)
}

// TagsGetter gets entries
type TagsGetter interface {
	GetAll() ([]models.Tag, []error)
}

// ReviewsGetter gets entries
type ReviewsGetter interface {
	GetAll() ([]models.Review, []error)
}
