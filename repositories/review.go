package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/webmalc/vishleva-backend/models"
)

// ReviewRepository is the repository.
type ReviewRepository struct {
	db *gorm.DB
}

// GetAll returns all entries.
func (r *ReviewRepository) GetAll() ([]models.Review, []error) {
	reviews := []models.Review{}
	r.db.Preload("Client").Preload("Image").
		Not("is_enabled", false).Find(&reviews)

	return reviews, r.db.GetErrors()
}

// NewReviewRepository returns a new repository struct.
func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}
