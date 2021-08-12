package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/models"
)

func TestCollectionRepository_GetAll(t *testing.T) {
	c := db.NewConnection()
	models.Migrate(c)
	repo := NewCollectionRepository(c.DB)
	collections, _ := repo.GetAll()
	collectionOne := &models.Collection{Name: "collection one", IsEnabled: true}
	collectionTwo := &models.Collection{Name: "collection two", IsEnabled: false}
	assert.Len(t, collections, 0)

	c.Create(collectionOne)
	c.Create(collectionTwo)
	collections, _ = repo.GetAll()
	assert.Len(t, collections, 1)
	assert.Equal(t, "collection one", collections[0].Name)
	assert.True(t, collections[0].IsEnabled)
}

func TestNewCollectionRepository(t *testing.T) {
	c := db.NewConnection()
	defer c.Close()
	r := NewCollectionRepository(c.DB)
	assert.Equal(t, r.db, c.DB)
}
