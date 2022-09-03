package services

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

// IsPositiveValidator checks if values is positive.
func IsPositiveValidator(i interface{}, name string, db *gorm.DB) {
	message := fmt.Sprintf("%s is negative", name)
	switch v := i.(type) {
	case int:
		if v < 0 {
			_ = db.AddError(errors.New(message))
		}
	case float64:
		if v < 0 {
			_ = db.AddError(errors.New(message))
		}
	case decimal.Decimal:
		if v.IsNegative() {
			_ = db.AddError(errors.New(message))
		}
	default:
		panic(fmt.Sprintf("unknown type %T!\n", v))
	}
}
