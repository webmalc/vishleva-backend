package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/handlers/mocks"
)

func TestCollectionHandler_GetList(t *testing.T) {
	checkResponse(t, "/api/collections", 1)
}

func TestNewCollectionHandler(t *testing.T) {
	rg := &mocks.CollectionsGetter{}
	handler := NewCollectionHandler(rg)
	assert.Equal(t, rg, handler.getter)
}
