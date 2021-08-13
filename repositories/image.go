package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/webmalc/vishleva-backend/models"
)

// ImageRepository is the repository
type ImageRepository struct {
	db *gorm.DB
}

// GetAll returns all entries
func (r *ImageRepository) GetAll(
	tag string,
	collectionID uint,
) ([]models.Image, []error) {
	images := []models.Image{}
	query := r.db.Preload("Tags")

	if collectionID > 0 || tag != "" {
		ids := []uint{}
		params := []interface{}{}
		raw := `SELECT it.image_id as id
			FROM image_tags as it
			JOIN tags as t ON it.tag_id = t.id`
		if tag != "" {
			raw += " WHERE t.name = ?"
			params = append(params, tag)
		}
		if collectionID > 0 {
			r := NewCollectionRepository(r.db)
			raw += " AND t.id IN (?)"
			params = append(params, r.GetTagsIDs(collectionID))
		}
		r.db.Raw(raw, params...).Pluck("it.image_id", &ids)
		query = query.Where("id IN (?)", ids)
	}
	query.Find(&images)
	return images, r.db.GetErrors()
}

// NewImageRepository returns a new repository struct
func NewImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{db: db}
}
