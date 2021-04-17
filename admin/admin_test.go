package admin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
)

// Should create a new admin
func TestNewAdmin(t *testing.T) {
	conn := db.NewConnection()
	adm := NewAdmin(conn.DB)
	assert.Equal(t, adm.db, conn.DB)
	assert.NotNil(t, adm.config)
	assert.NotNil(t, adm.admin)
}

// Should mount an admin
func TestAdmin_Mount(t *testing.T) {
	conn := db.NewConnection()
	adm := NewAdmin(conn.DB)
	mux := adm.Mount()
	assert.NotNil(t, mux)
}

func TestAdmin_GetBasePath(t *testing.T) {
	conn := db.NewConnection()
	adm := NewAdmin(conn.DB)
	assert.Equal(t, adm.GetBasePath(), "admin")
}
