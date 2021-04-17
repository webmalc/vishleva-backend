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
	"github.com/webmalc/vishleva-backend/admin"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/server/mocks"
)

// Should mount the admin
func TestServer_mountAdmin(t *testing.T) {
	l := &mocks.InfoLogger{}
	a := &mocks.Admin{}
	server := NewServer(a, l)
	server.engine = gin.Default()
	a.On("GetBasePath").Return("admin").Once()
	a.On("Mount").Return(http.NewServeMux()).Once()
	server.mountAdmin()
	l.AssertExpectations(t)
}

// Should init the logger
func TestServer_initLogger(t *testing.T) {
	assert.IsType(t, &os.File{}, gin.DefaultWriter)
	l := &mocks.InfoLogger{}
	a := &mocks.Admin{}
	server := NewServer(a, l)
	server.initLogger()
	assert.IsType(t, io.MultiWriter(), gin.DefaultWriter)
}

func TestServer_getCORS(t *testing.T) {
	l := &mocks.InfoLogger{}
	a := &mocks.Admin{}
	server := NewServer(a, l)
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
	a := &mocks.Admin{}
	server := NewServer(a, l)
	server.config.IsReleaseMode = true
	assert.Nil(t, server.engine)
	a.On("GetBasePath").Return("admin").Once()
	a.On("Mount").Return(http.NewServeMux()).Once()
	server.setEngine()
	l.AssertExpectations(t)
	assert.IsType(t, gin.Default(), server.engine)
	assert.Equal(t, gin.ReleaseMode, gin.Mode())
}

func TestServer_Run(t *testing.T) {
	l := &mocks.InfoLogger{}
	conn := db.NewConnection()
	a := admin.NewAdmin(conn.DB)
	server := NewServer(a, l)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go server.Run(ctx, []string{})
	time.Sleep(1 * time.Millisecond)
	resp, err := http.Get("http://localhost:9000/admin/")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	resp.Body.Close()
}

// Should create the server
func TestNewServer(t *testing.T) {
	l := &mocks.InfoLogger{}
	a := &mocks.Admin{}
	server := NewServer(a, l)
	assert.NotNil(t, server)
	assert.Equal(t, server.admin, a)
	assert.Equal(t, server.logger, l)
}
