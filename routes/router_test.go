package routes

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/routes/mocks"
)

func TestRouter_mountAdmin(t *testing.T) {
	h := &mocks.AuthHander{}
	ta := &mocks.ListHander{}
	tg := &mocks.ListHander{}
	rg := &mocks.ListHander{}
	ch := &mocks.ListHander{}
	ih := &mocks.ListHander{}
	a := &mocks.Admin{}
	r := NewRouter(a, h, ta, tg, rg, ch, ih)
	e := gin.Default()
	a.On("GetBasePath").Return("admin").Once()
	a.On("Mount").Return(http.NewServeMux()).Once()
	r.mountAdmin(e)
	a.AssertExpectations(t)
}

func TestRouter_BindRoutes(t *testing.T) {
	h := &mocks.AuthHander{}
	a := &mocks.Admin{}
	ta := &mocks.ListHander{}
	tg := &mocks.ListHander{}
	rg := &mocks.ListHander{}
	ch := &mocks.ListHander{}
	ih := &mocks.ListHander{}
	r := NewRouter(a, h, ta, tg, rg, ch, ih)
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
	ta := &mocks.ListHander{}
	tg := &mocks.ListHander{}
	rg := &mocks.ListHander{}
	ch := &mocks.ListHander{}
	ih := &mocks.ListHander{}
	r := NewRouter(a, h, ta, tg, rg, ch, ih)

	assert.Equal(t, r.auth, h)
	assert.Equal(t, r.admin, a)
	assert.Equal(t, r.tariffs, ta)
	assert.Equal(t, r.tags, tg)
	assert.NotNil(t, r.cacheStore)
}
