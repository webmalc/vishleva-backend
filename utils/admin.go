package utils

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
)

// GetDateFilter return an admin date filter.
func GetDateFilter(
	table, field string,
) func(*gorm.DB, *admin.FilterArgument) *gorm.DB {
	return func(db *gorm.DB, filterArgument *admin.FilterArgument) *gorm.DB {
		if d := filterArgument.Value.Get("Start"); d != nil {
			db = db.Where(
				fmt.Sprintf("%v.%v >= ?", table, field),
				d.Value,
			)
		}
		if d := filterArgument.Value.Get("End"); d != nil {
			db = db.Where(
				fmt.Sprintf("%v.%v <= ?", table, field),
				d.Value,
			)
		}

		return db
	}
}
