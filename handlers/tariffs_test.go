package handlers

import (
	"net/http"
	"testing"

	"github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/handlers/mocks"
)

func TestTariffsHander_GetList(t *testing.T) {
	w, engine := initRoutes()
	req, _ := http.NewRequest("GET", "/api/tariffs", nil)
	engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	json, err := simplejson.NewFromReader(w.Body)
	assert.Nil(t, err)
	entries, err := json.Get("entries").Array()
	assert.Nil(t, err)
	assert.Len(t, entries, 1)
}

func TestNewTariffsHandler(t *testing.T) {
	tg := &mocks.TariffsGetter{}
	handler := NewTariffsHandler(tg)
	assert.Equal(t, tg, handler.getter)
}
