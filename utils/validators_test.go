package utils

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
)

func TestIsPositiveValidator(t *testing.T) {
	conn := db.NewConnection()
	name := "test"
	assert.Empty(t, conn.GetErrors())

	IsPositiveValidator(1, name, conn.DB)
	assert.Empty(t, conn.GetErrors())

	IsPositiveValidator(0, name, conn.DB)
	assert.Empty(t, conn.GetErrors())

	IsPositiveValidator(-1, name, conn.DB)
	assert.Len(t, conn.GetErrors(), 1)
	assert.Contains(t, conn.GetErrors()[0].Error(), name)

	IsPositiveValidator(-1.3, name, conn.DB)
	assert.Len(t, conn.GetErrors(), 2)

	IsPositiveValidator(decimal.NewFromInt(-1), name, conn.DB)
	assert.Len(t, conn.GetErrors(), 3)
}
