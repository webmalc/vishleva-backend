package routes

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/routes/mocks"
)

func TestRouter_mountAdmin(t *testing.T) {
	h := &mocks.AuthHandler{}
	ta := &mocks.ListHandler{}
	tg := &mocks.ListHandler{}
	rg := &mocks.ListHandler{}
	ch := &mocks.ListHandler{}
	ih := &mocks.ListHandler{}
	ca := &mocks.ListHandler{}
	bo := &mocks.PostHandler{}
	a := &mocks.Admin{}
	r := NewRouter(a, h, ta, tg, rg, ch, ih, ca, bo)
	e := gin.Default()
	a.On("GetBasePath").Return("admin").Once()
	a.On("Mount").Return(http.NewServeMux()).Once()
	r.mountAdmin(e)
	a.AssertExpectations(t)
}

func TestRouter_BindRoutes(t *testing.T) {
	h := &mocks.AuthHandler{}
	a := &mocks.Admin{}
	ta := &mocks.ListHandler{}
	tg := &mocks.ListHandler{}
	rg := &mocks.ListHandler{}
	ch := &mocks.ListHandler{}
	ih := &mocks.ListHandler{}
	ca := &mocks.ListHandler{}
	bo := &mocks.PostHandler{}
	r := NewRouter(a, h, ta, tg, rg, ch, ih, ca, bo)
	e := gin.Default()
	a.On("GetBasePath").Return("admin").Once()
	a.On("Mount").Return(http.NewServeMux()).Once()
	r.BindRoutes(e)
	h.AssertExpectations(t)
	a.AssertExpectations(t)
}

func TestNewRouter(t *testing.T) {
	h := &mocks.AuthHandler{}
	a := &mocks.Admin{}
	ta := &mocks.ListHandler{}
	tg := &mocks.ListHandler{}
	rg := &mocks.ListHandler{}
	ch := &mocks.ListHandler{}
	ih := &mocks.ListHandler{}
	ca := &mocks.ListHandler{}
	bo := &mocks.PostHandler{}
	r := NewRouter(a, h, ta, tg, rg, ch, ih, ca, bo)

	assert.Equal(t, r.auth, h)
	assert.Equal(t, r.admin, a)
	assert.Equal(t, r.tariffs, ta)
	assert.Equal(t, r.tags, tg)
	assert.NotNil(t, r.cacheStore)
}
