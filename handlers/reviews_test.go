package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/handlers/mocks"
)

func TestReviewsHandler_GetList(t *testing.T) {
	checkResponse(t, "/api/reviews", 1)
}

func TestNewReviewsHandler(t *testing.T) {
	rg := &mocks.ReviewsGetter{}
	handler := NewReviewsHandler(rg)
	assert.Equal(t, rg, handler.getter)
}
