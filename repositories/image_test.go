package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/models"
)

func TestImageRepository_GetAll(t *testing.T) {
	c := db.NewConnection()
	models.Migrate(c)
	repo := NewImageRepository(c.DB)
	images, _ := repo.GetAll("", 0)
	imageOne := &models.Image{Name: "image one", Tags: []models.Tag{
		{Name: "tag one"},
		{Name: "tag two"},
	}}
	imageTwo := &models.Image{Name: "image two", Tags: []models.Tag{
		{Name: "tag three"},
		{Name: "tag four"},
	}}
	assert.Len(t, images, 0)

	c.Create(imageOne)
	c.Create(imageTwo)
	images, _ = repo.GetAll("", 0)
	assert.Len(t, images, 2)
	assert.Equal(t, "image one", images[0].Name)

	images, _ = repo.GetAll("tag four", 0)
	assert.Len(t, images, 1)
	assert.Equal(t, "image two", images[0].Name)
}

func TestNewImageRepository(t *testing.T) {
	c := db.NewConnection()
	defer c.Close()
	r := NewImageRepository(c.DB)
	assert.Equal(t, r.db, c.DB)
}
