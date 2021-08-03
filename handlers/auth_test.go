package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/admin"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/common/logger"
	"github.com/webmalc/vishleva-backend/common/session"
	"github.com/webmalc/vishleva-backend/handlers/mocks"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/repositories"
	"github.com/webmalc/vishleva-backend/routes"
)

func initRouters() (*httptest.ResponseRecorder, *gin.Engine) {
	log := logger.NewLogger()
	conn := db.NewConnection()
	sessionConfig := session.NewSession()
	userRepository := repositories.NewUserRepository(conn.DB)
	models.Migrate(conn)
	router := routes.NewRouter(
		admin.NewAdmin(conn.DB, sessionConfig),
		NewAuthHandler(sessionConfig, userRepository, log),
	)

	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*/*")
	engine.Use(sessions.Sessions(sessionConfig.Name, sessionConfig.Store))
	router.BindRoutes(engine)
	w := httptest.NewRecorder()
	return w, engine
}

func TestAuthHander_GetLogin(t *testing.T) {
	w, engine := initRouters()
	req, _ := http.NewRequest("GET", "/auth/login", nil)
	engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Sing in")
}

func TestAuthHander_PostLogin(t *testing.T) {
	w, engine := initRouters()
	req, _ := http.NewRequest("POST", "/auth/login", nil)
	engine.ServeHTTP(w, req)
	assert.Equal(t, 303, w.Code)
}

func TestAuthHander_GetLogout(t *testing.T) {
	w, engine := initRouters()
	req, _ := http.NewRequest("GET", "/auth/logout", nil)
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
