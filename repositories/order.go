package repositories

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"github.com/webmalc/vishleva-backend/models"
)

// OrderRepository is the repository.
type OrderRepository struct {
	db *gorm.DB
}

// CreateOnlineOrder creates online order.
func (r *OrderRepository) CreateOnlineOrder(
	name, comment string, begin, end *time.Time, client *models.Client,
) (*models.Order, error) {
	order := models.Order{
		Name:    name,
		Begin:   begin,
		End:     end,
		Comment: comment,
		Total:   decimal.NewFromInt(0),
		Paid:    decimal.NewFromInt(0),
		Client:  *client,
		Status:  "not_confirmed",
		Source:  "online",
	}
	err := r.db.Create(&order).Error

	return &order, err
}

// GetUpcoming returns all entries.
func (r *OrderRepository) GetUpcoming() ([]models.Order, []error) {
	orders := []models.Order{}
	d := time.Now()
	m := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local)
	r.db.Where(`"end" > ?`, m).
		Not("status", []string{"not_confirmed", "closed"}).
		Find(&orders)

	return orders, r.db.GetErrors()
}

// NewOrderRepository returns a new repository struct.
func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}
