package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/handlers/mocks"
)

func TestTagsHander_GetList(t *testing.T) {
	checkResponse(t, "/api/tags", 2)
}

func TestNewTagsHandler(t *testing.T) {
	tg := &mocks.TagsGetter{}
	handler := NewTagsHandler(tg)
	assert.Equal(t, tg, handler.getter)
}
