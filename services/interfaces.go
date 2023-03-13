package services

import (
	"time"

	"github.com/webmalc/vishleva-backend/models"
)

// InfoLogger logs errors.
type InfoLogger interface {
	Infof(format string, args ...interface{})
}

// ClientGetterCreator gets or create a client.
type ClientGetterCreator interface {
	GetOrCreate(email, phone, name string) (*models.Client, error)
}

// OrderBooker crete a new book order.
type OrderBooker interface {
	CreateOnlineOrder(
		name, comment string, begin, end *time.Time, client *models.Client,
	) (*models.Order, error)
}

// AdminNotifier notify admin.
type AdminNotifier interface {
	notifyAdmin(subject string, data map[string]string) error
}
