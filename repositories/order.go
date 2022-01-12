package repositories

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/webmalc/vishleva-backend/models"
)

// OrderRepository is the repository
type OrderRepository struct {
	db *gorm.DB
}

// GetAll returns all entries
func (r *OrderRepository) GetUpcoming() ([]models.Order, []error) {
	orders := []models.Order{}
	d := time.Now()
	m := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local)
	r.db.Where(`"end" > ?`, m).
		Not("status", []string{"not_confirmed", "closed"}).
		Find(&orders)
	return orders, r.db.GetErrors()
}

// NewOrderRepository returns a new repository struct
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}
