package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/models"
)

func TestTagRepository_GetAll(t *testing.T) {
	c := db.NewConnection()
	models.Migrate(c)
	repo := NewTagRepository(c.DB)
	tags, _ := repo.GetAll()
	tag := &models.Tag{
		Name: "one",
	}
	assert.Len(t, tags, 0)

	c.Create(tag)
	tags, _ = repo.GetAll()
	assert.Len(t, tags, 1)
	assert.Equal(t, "one", tags[0].Name)
}

func TestNewTagRepository(t *testing.T) {
	c := db.NewConnection()
	defer c.Close()
	r := NewTagRepository(c.DB)
	assert.Equal(t, r.db, c.DB)
}
