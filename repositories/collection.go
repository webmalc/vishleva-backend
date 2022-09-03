package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/webmalc/vishleva-backend/models"
)

// CollectionRepository is the repository.
type CollectionRepository struct {
	db *gorm.DB
}

// GetAll returns all entries.
func (r *CollectionRepository) GetAll() ([]models.Collection, []error) {
	collections := []models.Collection{}
	r.db.Preload("Image").Not("is_enabled", false).
		Find(&collections)

	return collections, r.db.GetErrors()
}

// GetTagsIDs return the collection tags IDs.
func (r *CollectionRepository) GetTagsIDs(collectionID uint) []uint {
	ids := []uint{}
	r.db.Raw(`
	SELECT ct.tag_id FROM collection_tags as ct
	WHERE ct.collection_id = ?`, collectionID).
		Pluck("", &ids)

	return ids
}

// NewCollectionRepository returns a new repository struct.
func NewCollectionRepository(db *gorm.DB) *CollectionRepository {
	return &CollectionRepository{db: db}
}
