package utils

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStringInSlice(t *testing.T) {
	i, b := StringInSlice("bar", []string{"spam", "foo", "bar"})
	assert.True(t, b)
	assert.Equal(t, 2, i)

	i, b = StringInSlice("invalid", []string{"spam", "foo", "bar"})
	assert.False(t, b)
	assert.Equal(t, -1, i)
}

func TestStructToMap(t *testing.T) {
	now := time.Now()
	r := StructToMap(struct {
		Name  string
		Order int
		Begin time.Time
	}{Name: "test name", Order: 9, Begin: now})
	assert.Equal(t, "test name", r["Name"])
	assert.Equal(t, "9", r["Order"])
	assert.Equal(t, now.Format("2 January 15:04"), r["Begin"])
	assert.Contains(t, r["Begin"], strconv.Itoa(now.Day()))
}
