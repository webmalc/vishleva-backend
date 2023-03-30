package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/dto"
	"github.com/webmalc/vishleva-backend/handlers/mocks"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/repositories"
	"github.com/webmalc/vishleva-backend/services"
)

func TestBookHandler_Post(t *testing.T) {
	w, engine := initRoutes()
	now := time.Now()
	begin := time.Date(
		now.Year()+1, now.Month(), now.Day(), 16, 0, 0, 0, time.Local,
	)
	end := begin.Add(time.Hour)
	bookDto := &dto.Book{
		Name:       "name",
		Comment:    "comment comment",
		Begin:      begin,
		End:        end,
		Phone:      "79251112233",
		ClientName: "client",
		Email:      "test@test.com",
	}
	jsonStr, _ := json.Marshal(bookDto) // nolint
	req, _ := http.NewRequest("POST", "/api/book", bytes.NewBuffer(jsonStr))
	engine.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}

func TestNewBookHandler(t *testing.T) {
	log := &mocks.InfoLogger{}
	conn := db.NewConnection()
	models.Migrate(conn)
	orderRepository := repositories.NewOrderRepository(conn.DB)
	clientRepository := repositories.NewClientRepository(conn.DB)
	service := services.NewBookingService(
		log,
		clientRepository,
		orderRepository,
	)
	handler := NewBookHandler(service)
	assert.Equal(t, service, handler.booker)
}
