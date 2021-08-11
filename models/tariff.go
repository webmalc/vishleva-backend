package models

import (
	"encoding/json"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/qor/sorting"
	"github.com/shopspring/decimal"
	"github.com/webmalc/vishleva-backend/services"
)

// Tariff is a model
type Tariff struct {
	gorm.Model
	sorting.Sorting
	Name         string          `gorm:"size:255;not null;index;unique" valid:"required"`
	Price        decimal.Decimal `sql:"type:decimal(20,2);" gorm:";not null;" valid:"required"`
	Duration     int             `gorm:";not null;" valid:"required"`
	Photos       int             `gorm:";not null;" valid:"required;duck"`
	Retouch      int             `gorm:";not null;" valid:"required"`
	RetouchPrice decimal.Decimal `sql:"type:decimal(20,2);" gorm:";not null;" valid:"required"`
	Description  string          `gorm:"type:text; index" json:"description"`
	IsEnabled    bool            `gorm:"type:bool;default:false;index"`
	IsPrimary    bool            `gorm:"type:bool;default:false;index"`
}

// Validate validates the struct
func (t *Tariff) Validate(db *gorm.DB) {
	services.IsPositiveValidator(t.Price, "price", db)
	services.IsPositiveValidator(t.Duration, "duration", db)
	services.IsPositiveValidator(t.Photos, "photos", db)
	services.IsPositiveValidator(t.Retouch, "retouch", db)
	services.IsPositiveValidator(t.RetouchPrice, "retouch price", db)
	if t.Retouch > t.Photos {
		_ = db.AddError(errors.New(
			"retouch number is greater than the number of photos",
		))
	}
}

// MarshalJSON returns the JSON respresentation
func (t *Tariff) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		Price        string `json:"price"`
		Duration     int    `json:"duration"`
		Photos       int    `json:"photos"`
		Retouch      int    `json:"retouch"`
		RetouchPrice string `json:"retouch_price"`
		IsPrimary    bool   `json:"is_primary"`
	}{
		Name:         t.Name,
		Description:  t.Description,
		Price:        t.Price.String(),
		Duration:     t.Duration,
		Photos:       t.Photos,
		Retouch:      t.Retouch,
		RetouchPrice: t.RetouchPrice.String(),
		IsPrimary:    t.IsPrimary,
	})
}
