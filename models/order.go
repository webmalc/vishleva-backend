package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"github.com/webmalc/vishleva-backend/utils"
)

// Order is a model.
type Order struct {
	gorm.Model
	Name     string          `gorm:"size:255;not null;index" valid:"required"`
	Begin    *time.Time      `gorm:"index;not null" valid:"required"`
	End      *time.Time      `gorm:"index;not null" valid:"required"`
	Comment  string          `gorm:"type:text;index"`
	Total    decimal.Decimal `sql:"type:decimal(20,2);" gorm:";not null;" valid:"required"`
	Paid     decimal.Decimal `sql:"type:decimal(20,2);" gorm:";not null;" valid:"required"`
	ClientID uint            `gorm:"type:bigint;index;not null"`
	Client   Client          `gorm:"constraint:OnDelete:RESTRICT;not null"`
	Status   string          `gorm:"index;not null;default:'not_confirmed'" valid:"required"`
	Source   string          `gorm:"index;not null;default:'manual'" valid:"required"`
}

// Validate validates the client.
func (t *Order) Validate(db *gorm.DB) {
	c := NewConfig()
	// check dates
	if t.Begin.After(*t.End) || t.Begin.Equal(*t.End) {
		_ = db.AddError(errors.New(
			"the begin is equal or greater than the end",
		))
	}

	// check dates for online orders
	if t.Source == "online" {
		utils.IsDateInFutureValidator(*t.Begin, "begin", db)
		utils.IsDateInFutureValidator(*t.End, "end", db)
	}

	// check price
	utils.IsPositiveValidator(t.Total, "total", db)
	utils.IsPositiveValidator(t.Paid, "paid", db)

	// status and source
	if _, ok := utils.StringInSlice(t.Status, c.OrderStatuses); !ok {
		_ = db.AddError(errors.New("status is invalid"))
	}
	if _, ok := utils.StringInSlice(t.Source, c.OrderSources); !ok {
		_ = db.AddError(errors.New("source is invalid"))
	}

	// check overlapping
	if t.Status != "open" {
		return
	}
	count := 0
	db.Where(`begin <= ? AND "end" >= ?`, *t.End, *t.Begin).
		Not("status", []string{"not_confirmed", "closed"}).
		Not("id", t.ID).
		Find(&[]Order{}).Count(&count)
	if count > 0 {
		_ = db.AddError(errors.New("the order is overlapping"))
	}
}
