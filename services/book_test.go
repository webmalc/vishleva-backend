package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/common/test"
	"github.com/webmalc/vishleva-backend/dto"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/repositories"
	"github.com/webmalc/vishleva-backend/services/mocks"
)

func TestBookingService_Book(t *testing.T) {
	log := &mocks.InfoLogger{}
	conn := db.NewConnection()
	orders := []models.Order{}
	clients := []models.Client{}
	models.Migrate(conn)
	orderRepository := repositories.NewOrderRepository(conn.DB)
	clientRepository := repositories.NewClientRepository(conn.DB)
	service := NewBookingService(
		log,
		clientRepository,
		orderRepository,
	)
	now := time.Now()
	begin := time.Date(
		now.Year()+1, now.Month(), now.Day(), 16, 0, 0, 0, time.Local,
	)
	end := begin.Add(time.Hour)

	bookDto := &dto.Book{
		Email:      "t@test.com",
		Phone:      "79251112233",
		ClientName: "client",
		Name:       "test name",
		Comment:    "test comment",
		Begin:      begin,
		End:        end,
	}
	log.On("Infof", mock.Anything, bookDto).Return(nil).Twice()

	conn.DB.Find(&orders)
	conn.DB.Find(&clients)
	assert.Len(t, orders, 0)
	assert.Len(t, clients, 0)

	// // book the first order
	order, _ := service.Book(bookDto)
	conn.DB.Find(&orders)
	conn.DB.Find(&clients)

	assert.NotNil(t, order)
	assert.Greater(t, order.ID, uint(0))
	assert.Equal(t, "not_confirmed", order.Status)
	assert.Equal(t, "online", order.Source)
	assert.Equal(t, "t@test.com", *order.Client.Email)
	assert.Len(t, orders, 1)
	assert.Len(t, clients, 1)

	// book the second order
	_, _ = service.Book(bookDto)
	conn.DB.Find(&orders)
	conn.DB.Find(&clients)
	assert.Len(t, orders, 2)
	assert.Len(t, clients, 1)
	log.AssertExpectations(t)
}

func TestNewBookingService(t *testing.T) {
	log := &mocks.InfoLogger{}
	conn := db.NewConnection()
	orderRepository := repositories.NewOrderRepository(conn.DB)
	clientRepository := repositories.NewClientRepository(conn.DB)
	service := NewBookingService(
		log,
		clientRepository,
		orderRepository,
	)
	assert.Equal(t, orderRepository, service.orderBooker)
	assert.Equal(t, clientRepository, service.clientGetterCreator)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
