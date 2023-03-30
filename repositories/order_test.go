package repositories

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/models"
)

func TestOrderRepository_CreateOnlineOrder(t *testing.T) {
	conn := db.NewConnection()
	models.Migrate(conn)
	repo := NewOrderRepository(conn.DB)
	clientRepo := NewClientRepository(conn.DB)
	now := time.Now()
	begin := time.Date(
		now.Year()+1, now.Month(), now.Day(), 16, 0, 0, 0, time.Local,
	)
	end := begin.Add(time.Hour)
	client, _ := clientRepo.GetOrCreate("t@test.com", "79251112233", "client")
	order, _ := repo.CreateOnlineOrder(
		"new order", "comment", &begin, &end, client,
	)
	assert.Greater(t, order.ID, uint(0))
	assert.Equal(t, "not_confirmed", order.Status)
	assert.Equal(t, "online", order.Source)
}

func TestOrderRepository_GetUpcoming(t *testing.T) {
	conn := db.NewConnection()
	models.Migrate(conn)
	repo := NewOrderRepository(conn.DB)
	result, _ := repo.GetUpcoming()
	assert.Empty(t, result)

	now := time.Now()
	begin := time.Date(
		now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, time.Local,
	)
	end := begin.Add(time.Hour)
	conn.Create(&models.Order{
		Name:   "test order 1",
		Begin:  &begin,
		End:    &end,
		Total:  decimal.NewFromInt(10),
		Paid:   decimal.NewFromInt(10),
		Status: "open",
		Source: "manual",
	})
	conn.Create(&models.Order{
		Name:   "test order 2",
		Begin:  &begin,
		End:    &end,
		Total:  decimal.NewFromInt(10),
		Paid:   decimal.NewFromInt(10),
		Status: "closed",
		Source: "manual",
	})

	result, _ = repo.GetUpcoming()
	assert.Len(t, result, 1)
	assert.Equal(t, "test order 1", result[0].Name)
}

func TestNewOrderRepository(t *testing.T) {
	conn := db.NewConnection()
	repo := NewOrderRepository(conn.DB)
	assert.Equal(t, repo.db, conn.DB)
}
