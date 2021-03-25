package db

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// Should create a new database connection.
func TestNewConnection(t *testing.T) {
	conn := NewConnection()
	defer conn.Close()
	assert.NotNil(t, conn)
	err := conn.DB.DB().Ping()
	assert.NoError(t, err)
}

// Should panic.
func TestNewLoggerPanic(t *testing.T) {
	o := viper.Get(databaseKey)
	defer viper.Set(databaseKey, o)

	viper.Set(databaseKey, "/invalid/path")
	assert.Panics(t, func() {
		NewConnection()
	})
}
