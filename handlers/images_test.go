package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/handlers/mocks"
)

func TestImagesHander_GetList(t *testing.T) {
	checkResponse(t, "/api/images", 2)
}

func TestNewImagesHandler(t *testing.T) {
	ig := &mocks.ImagesGetter{}
	handler := NewImagesHandler(ig)
	assert.Equal(t, ig, handler.getter)
}
