package repositories

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/models"
)

func TestTariffRepository_GetAll(t *testing.T) {
	c := db.NewConnection()
	models.Migrate(c)
	repo := NewTariffRepository(c.DB)
	tariffs, _ := repo.GetAll()
	tariffOne := &models.Tariff{
		Name:         "one",
		Price:        decimal.NewFromInt(10),
		Duration:     10,
		Photos:       20,
		Retouch:      10,
		RetouchPrice: decimal.NewFromInt(10),
		IsEnabled:    true,
	}
	tariffTwo := &models.Tariff{
		Name:         "two",
		Price:        decimal.NewFromInt(10),
		Duration:     10,
		Photos:       20,
		Retouch:      10,
		RetouchPrice: decimal.NewFromInt(10),
		IsEnabled:    false,
	}
	assert.Len(t, tariffs, 0)

	c.Create(tariffOne)
	c.Create(tariffTwo)
	tariffs, _ = repo.GetAll()
	assert.Len(t, tariffs, 1)
	assert.Equal(t, "one", tariffs[0].Name)
	assert.True(t, tariffs[0].IsEnabled)
}

func TestNewTariffRepository(t *testing.T) {
	c := db.NewConnection()
	defer c.Close()
	r := NewTariffRepository(c.DB)
	assert.Equal(t, r.db, c.DB)
}
