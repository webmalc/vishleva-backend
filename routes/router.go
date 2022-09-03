package routes

import (
	"fmt"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Router is the router
type Router struct {
	admin       Admin
	config      *Config
	auth        AuthHandler
	tariffs     ListHandler
	tags        ListHandler
	reviews     ListHandler
	collections ListHandler
	images      ListHandler
	calendar    ListHandler
	cacheStore  *persistence.InMemoryStore
}

// mountAdmin mount the admin
func (r *Router) mountAdmin(e *gin.Engine) {
	e.Any(
		fmt.Sprintf("/%s/*resources", r.admin.GetBasePath()),
		gin.WrapH(r.admin.Mount()),
	)
}

// BindRoutes binds the routes
func (r *Router) BindRoutes(e *gin.Engine) {
	r.mountAdmin(e)

	// static
	e.Use(static.Serve("/", static.LocalFile("./public", false)))

	// auth routes
	auth := e.Group("auth")
	auth.GET("/login", r.auth.GetLogin)
	auth.POST("/login", r.auth.PostLogin)
	auth.GET("/logout", r.auth.GetLogout)

	// api routes
	api := e.Group("api")
	api.GET("/tariffs", cache.CachePage(
		r.cacheStore, r.config.CacheTimeout, r.tariffs.GetList,
	))
	api.GET("/tags", cache.CachePage(
		r.cacheStore, r.config.CacheTimeout, r.tags.GetList,
	))
	api.GET("/reviews", cache.CachePage(
		r.cacheStore, r.config.CacheTimeout, r.reviews.GetList,
	))
	api.GET("/collections", cache.CachePage(
		r.cacheStore, r.config.CacheTimeout, r.collections.GetList,
	))
	api.GET("/images", cache.CachePage(
		r.cacheStore, r.config.CacheTimeout, r.images.GetList,
	))
	api.GET("/calendar", r.calendar.GetList)

	// cache
	api.GET("/cache", func(c *gin.Context) {
		r.cacheStore.Flush()
	})
}

// NewRouter returns a new router object
func NewRouter(
	admin Admin,
	auth AuthHandler,
	tariffs ListHandler,
	tags ListHandler,
	reviews ListHandler,
	collections ListHandler,
	images ListHandler,
	calendar ListHandler,
) *Router {
	config := NewConfig()
	return &Router{
		config:      config,
		admin:       admin,
		auth:        auth,
		tariffs:     tariffs,
		tags:        tags,
		reviews:     reviews,
		collections: collections,
		images:      images,
		calendar:    calendar,
		cacheStore:  persistence.NewInMemoryStore(config.CacheTimeout),
	}
}
