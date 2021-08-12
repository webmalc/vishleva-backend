package server

import (
	"context"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/vishleva-backend/admin"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/common/session"
	"github.com/webmalc/vishleva-backend/handlers"
	"github.com/webmalc/vishleva-backend/repositories"
	"github.com/webmalc/vishleva-backend/routes"
	"github.com/webmalc/vishleva-backend/server/mocks"
)

// Should init the logger
func TestServer_initLogger(t *testing.T) {
	assert.IsType(t, &os.File{}, gin.DefaultWriter)
	l := &mocks.InfoLogger{}
	r := &mocks.Router{}
	s := session.NewSession()
	server := NewServer(r, l, s)
	server.initLogger()
	assert.IsType(t, io.MultiWriter(), gin.DefaultWriter)
}

func TestServer_getCORS(t *testing.T) {
	l := &mocks.InfoLogger{}
	r := &mocks.Router{}
	s := session.NewSession()
	server := NewServer(r, l, s)
	cors := server.getCORS()
	assert.True(t, cors.AllowAllOrigins)
	assert.Contains(t, cors.AllowHeaders, "Authorization")

	server.config.IsReleaseMode = true
	server.config.ServerAllowOrigins = []string{"test_origin"}
	cors = server.getCORS()
	assert.False(t, cors.AllowAllOrigins)
	assert.Contains(t, cors.AllowOrigins, "test_origin")
}

func TestServer_setEngine(t *testing.T) {
	l := &mocks.InfoLogger{}
	r := &mocks.Router{}
	s := session.NewSession()
	server := NewServer(r, l, s)
	server.config.IsReleaseMode = true
	assert.Nil(t, server.engine)
	r.On("BindRoutes", mock.Anything).Once()
	server.setEngine()
	l.AssertExpectations(t)
	assert.IsType(t, gin.Default(), server.engine)
	assert.Equal(t, gin.ReleaseMode, gin.Mode())
}

func TestServer_Run(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	l := &mocks.InfoLogger{}
	e := &mocks.ErrorLogger{}
	conn := db.NewConnection()
	defer conn.Close()
	ur := repositories.NewUserRepository(conn.DB)
	tr := repositories.NewTariffRepository(conn.DB)
	tg := repositories.NewTagRepository(conn.DB)
	rr := repositories.NewReviewRepository(conn.DB)
	s := session.NewSession()
	a := admin.NewAdmin(conn.DB, s)
	r := routes.NewRouter(
		a,
		handlers.NewAuthHandler(s, ur, e),
		handlers.NewTariffsHandler(tr),
		handlers.NewTagsHandler(tg),
		handlers.NewReviewsHandler(rr),
	)
	server := NewServer(r, l, s)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go server.Run(ctx, []string{})
	time.Sleep(1 * time.Millisecond)
	resp, err := http.Get("http://localhost:9000/auth/login/")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp.Body.Close()
}

// Should create the server
func TestNewServer(t *testing.T) {
	l := &mocks.InfoLogger{}
	r := &mocks.Router{}
	s := session.NewSession()
	server := NewServer(r, l, s)
	assert.NotNil(t, server)
	assert.Equal(t, server.router, r)
	assert.Equal(t, server.logger, l)
}
