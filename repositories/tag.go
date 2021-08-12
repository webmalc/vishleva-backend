package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/webmalc/vishleva-backend/models"
)

// TagRepository is the repository
type TagRepository struct {
	db *gorm.DB
}

// GetAll returns all entries
func (r *TagRepository) GetAll() ([]models.Tag, []error) {
	tags := []models.Tag{}
	r.db.Find(&tags)

	return tags, r.db.GetErrors()
}

// NewTagRepository returns a new repository struct
func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}
