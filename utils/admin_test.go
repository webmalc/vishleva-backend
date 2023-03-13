package utils

import (
	"testing"

	"github.com/qor/admin"
	"github.com/qor/qor/resource"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
)

func TestGetDateFilter(t *testing.T) {
	conn := db.NewConnection()
	filterFunc := GetDateFilter("user", "created_at")
	filter := &admin.FilterArgument{Value: &resource.MetaValues{
		Values: []*resource.MetaValue{{
			Name:  "Start",
			Value: "start",
		}, {
			Name:  "End",
			Value: "end",
		}},
	}}
	assert.NotNil(t, filterFunc)
	d := filterFunc(conn.DB, filter)
	assert.NotNil(t, d)
}
