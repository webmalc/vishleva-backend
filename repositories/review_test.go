package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/models"
)

func TestReviewRepository_GetAll(t *testing.T) {
	c := db.NewConnection()
	models.Migrate(c)
	repo := NewReviewRepository(c.DB)
	reviews, _ := repo.GetAll()
	reviewOne := &models.Review{Content: "one", IsEnabled: true}
	reviewTwo := &models.Review{Content: "two", IsEnabled: false}
	assert.Len(t, reviews, 0)

	c.Create(reviewOne)
	c.Create(reviewTwo)
	reviews, _ = repo.GetAll()
	assert.Len(t, reviews, 1)
	assert.Equal(t, "one", reviews[0].Content)
	assert.True(t, reviews[0].IsEnabled)
}

func TestNewReviewRepository(t *testing.T) {
	c := db.NewConnection()
	defer c.Close()
	r := NewReviewRepository(c.DB)
	assert.Equal(t, r.db, c.DB)
}
