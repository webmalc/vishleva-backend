package handlers

import (
	"time"

	"github.com/webmalc/vishleva-backend/dto"
	"github.com/webmalc/vishleva-backend/models"
)

// ErrorLogger logs errors.
type ErrorLogger interface {
	Errorf(format string, args ...interface{})
}

// UserLoginer logs users.
type UserLoginer interface {
	LoginAndReturnUser(email, password string) (*models.User, error)
}

// TariffsGetter gets entries.
type TariffsGetter interface {
	GetAll() ([]models.Tariff, []error)
}

// TagsGetter gets entries.
type TagsGetter interface {
	GetAll() ([]models.Tag, []error)
}

// ReviewsGetter gets entries.
type ReviewsGetter interface {
	GetAll() ([]models.Review, []error)
}

// CollectionsGetter gets entries.
type CollectionsGetter interface {
	GetAll() ([]models.Collection, []error)
}

// ImagesGetter gets entries.
type ImagesGetter interface {
	GetAll(tag string, collectionID uint) ([]models.Image, []error)
}

// CalendarGenerator gets entries.
type CalendarGenerator interface {
	Get(begin time.Time) []*models.CalendarDay
}

// Booker books orders.
type Booker interface {
	Book(*dto.Book) (*models.Order, error)
}
