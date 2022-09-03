package handlers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/session"
	"github.com/webmalc/vishleva-backend/handlers/mocks"
)

func TestAuthHandler_GetLogin(t *testing.T) {
	w, engine := initRoutes()
	req, _ := http.NewRequest("GET", "/auth/login", http.NoBody)
	engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Sing in")
}

func TestAuthHandler_PostLogin(t *testing.T) {
	w, engine := initRoutes()
	req, _ := http.NewRequest("POST", "/auth/login", http.NoBody)
	engine.ServeHTTP(w, req)
	assert.Equal(t, 303, w.Code)
}

func TestAuthHandler_GetLogout(t *testing.T) {
	w, engine := initRoutes()
	req, _ := http.NewRequest("GET", "/auth/logout", http.NoBody)
	engine.ServeHTTP(w, req)
	assert.Equal(t, 303, w.Code)
}

func TestNewAuthHandler(t *testing.T) {
	log := &mocks.ErrorLogger{}
	ses := session.NewSession()
	userLoginer := &mocks.UserLogger{}
	handler := NewAuthHandler(ses, userLoginer, log)
	assert.Equal(t, log, handler.logger)
	assert.Equal(t, ses, handler.session)
	assert.Equal(t, userLoginer, handler.userLoginer)
}
