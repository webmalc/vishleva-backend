package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/handlers/mocks"
)

func TestTariffsHander_GetList(t *testing.T) {
	checkResponse(t, "/api/tariffs", 1)
}

func TestNewTariffsHandler(t *testing.T) {
	tg := &mocks.TariffsGetter{}
	handler := NewTariffsHandler(tg)
	assert.Equal(t, tg, handler.getter)
}
