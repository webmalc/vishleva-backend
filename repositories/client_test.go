package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/models"
)

func TestClientRepository_GetOrCreate(t *testing.T) {
	conn := db.NewConnection()
	models.Migrate(conn)
	repo := NewClientRepository(conn.DB)
	client, _ := repo.GetOrCreate("t@test.com", "79251112233", "client")

	assert.Greater(t, client.ID, uint(0))
	assert.Equal(t, "t@test.com", *client.Email)
	assert.Equal(t, "79251112233", client.Phone)
	assert.Equal(t, "automatically created client", client.Comment)
}

func TestNewClientRepository(t *testing.T) {
	conn := db.NewConnection()
	repo := NewClientRepository(conn.DB)
	assert.Equal(t, repo.db, conn.DB)
}
