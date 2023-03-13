package services

import (
	"github.com/webmalc/vishleva-backend/dto"
	"github.com/webmalc/vishleva-backend/models"

	"github.com/pkg/errors"
)

// BookingService is the booking service.
type BookingService struct {
	logger              InfoLogger
	clientGetterCreator ClientGetterCreator
	orderBooker         OrderBooker
	// notifier            AdminNotifier
}

// Book book orders.
func (s *BookingService) Book(data *dto.Book) (*models.Order, error) {
	s.logger.Infof("Book request has been received %+v", data)
	client, err := s.clientGetterCreator.GetOrCreate(
		data.Email, data.Phone, data.ClientName,
	)
	if err != nil {
		return nil, errors.Wrap(err, "booking")
	}
	order, err := s.orderBooker.CreateOnlineOrder(
		data.Name, data.Comment, &data.Begin, &data.End, client,
	)
	if err != nil {
		return nil, errors.Wrap(err, "booking")
	}

	return order, err
}

// NewBookingService returns a new booking service.
func NewBookingService(
	logger InfoLogger, client ClientGetterCreator, order OrderBooker,
) *BookingService {
	return &BookingService{
		logger:              logger,
		clientGetterCreator: client,
		orderBooker:         order,
	}
}
