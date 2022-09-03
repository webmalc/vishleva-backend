package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/webmalc/vishleva-backend/models"
)

// TariffRepository is the repository.
type TariffRepository struct {
	db *gorm.DB
}

// GetAll returns all entries.
func (r *TariffRepository) GetAll() ([]models.Tariff, []error) {
	tariffs := []models.Tariff{}
	r.db.Not("is_enabled", false).Find(&tariffs)

	return tariffs, r.db.GetErrors()
}

// NewTariffRepository returns a new repository struct.
func NewTariffRepository(db *gorm.DB) *TariffRepository {
	return &TariffRepository{db: db}
}
