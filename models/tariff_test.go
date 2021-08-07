package models

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
)

func TestTariff_Validate(t *testing.T) {
	conn := db.NewConnection()
	tariff := &Tariff{
		Name:         "test",
		Price:        decimal.NewFromInt(10),
		Duration:     10,
		Photos:       20,
		Retouch:      10,
		RetouchPrice: decimal.NewFromInt(10),
	}
	tariff.Validate(conn.DB)
	assert.Len(t, conn.GetErrors(), 0)

	tariff.Retouch = 100
	tariff.Validate(conn.DB)
	assert.Len(t, conn.GetErrors(), 1)
	assert.Contains(
		t, conn.GetErrors()[0].Error(), "retouch number is greater",
	)
}
