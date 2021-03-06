package models

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
)

func TestOrder_Validate(t *testing.T) {
	conn := db.NewConnection()
	begin := time.Now()
	end := begin.Add(time.Hour)
	order := &Order{
		Name:   "test",
		Begin:  &begin,
		End:    &end,
		Total:  decimal.NewFromInt(100),
		Paid:   decimal.NewFromInt(50),
		Status: "open",
	}
	order.Validate(conn.DB)
	assert.Len(t, conn.GetErrors(), 0)

	invalidEnd := end.Add(-time.Hour * 10)
	order.End = &invalidEnd

	order.Validate(conn.DB)
	assert.Len(t, conn.GetErrors(), 1)
	assert.Contains(
		t, conn.GetErrors()[0].Error(), "greater than the end",
	)
	order.End = &end

	order.Total = decimal.NewFromInt(-100)
	order.Validate(conn.DB)
	assert.Len(t, conn.GetErrors(), 2)
	assert.Contains(
		t, conn.GetErrors()[1].Error(), "total",
	)
	order.Total = decimal.NewFromInt(100)

	order.Paid = decimal.NewFromInt(-100)
	order.Validate(conn.DB)
	assert.Len(t, conn.GetErrors(), 3)
	assert.Contains(
		t, conn.GetErrors()[2].Error(), "paid",
	)
	order.Paid = decimal.NewFromInt(50)

	order.Status = "invalid"
	order.Validate(conn.DB)
	assert.Len(t, conn.GetErrors(), 4)
	assert.Contains(
		t, conn.GetErrors()[3].Error(), "invalid",
	)
}

func TestOrder_ValidateOverlapping(t *testing.T) {
	conn := db.NewConnection()
	conn.AutoMigrate(&Order{})
	begin := time.Now()
	end := begin.Add(time.Hour)
	orderFirst := &Order{
		Name:   "test",
		Begin:  &begin,
		End:    &end,
		Total:  decimal.NewFromInt(100),
		Paid:   decimal.NewFromInt(50),
		Status: "open",
	}

	orderFirst.Validate(conn.DB)
	conn.Create(&orderFirst)
	assert.Len(t, conn.GetErrors(), 0)

	count := 0

	conn.Find(&[]Order{}).Count(&count)
	assert.Equal(t, count, 1)

	orderSecond := &Order{
		Name:   "test",
		Begin:  &begin,
		End:    &end,
		Total:  decimal.NewFromInt(100),
		Paid:   decimal.NewFromInt(50),
		Status: "open",
	}

	orderSecond.Validate(conn.DB)
	assert.Len(t, conn.GetErrors(), 1)

	begin = begin.Add(time.Minute * 10)
	end = begin.Add(time.Minute * 10)
	orderSecond.Begin = &begin
	orderSecond.End = &end

	orderSecond.Validate(conn.DB)
	assert.Len(t, conn.GetErrors(), 2)
	assert.Contains(
		t, conn.GetErrors()[1].Error(), "is overlapping",
	)

	begin = begin.Add(time.Hour * 10)
	end = begin.Add(time.Hour * 10)

	orderSecond.Validate(conn.DB)
	assert.Len(t, conn.GetErrors(), 2)
}
