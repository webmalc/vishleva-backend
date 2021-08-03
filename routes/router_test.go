package routes

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/test"
	"github.com/webmalc/vishleva-backend/routes/mocks"
)

func TestRouter_mountAdmin(t *testing.T) {
	h := &mocks.AuthHander{}
	a := &mocks.Admin{}
	r := NewRouter(a, h)
	e := gin.Default()
	a.On("GetBasePath").Return("admin").Once()
	a.On("Mount").Return(http.NewServeMux()).Once()
	r.mountAdmin(e)
	a.AssertExpectations(t)
}

func TestRouter_BindRoutes(t *testing.T) {
	h := &mocks.AuthHander{}
	a := &mocks.Admin{}
	r := NewRouter(a, h)
	e := gin.Default()
	a.On("GetBasePath").Return("admin").Once()
	a.On("Mount").Return(http.NewServeMux()).Once()
	r.BindRoutes(e)
	h.AssertExpectations(t)
	a.AssertExpectations(t)
}

func TestNewRouter(t *testing.T) {
	h := &mocks.AuthHander{}
	a := &mocks.Admin{}
	r := NewRouter(a, h)

	assert.Equal(t, r.auth, h)
	assert.Equal(t, r.admin, a)
}

// Setups the tests
func TestMain(m *testing.M) {
	test.Run(m)
}
