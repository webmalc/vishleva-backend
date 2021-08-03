package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Router is the router
type Router struct {
	admin Admin
	auth  AuthHander
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

	// auth routes
	a := e.Group("auth")
	a.GET("/login", r.auth.GetLogin)
	a.POST("/login", r.auth.PostLogin)
	a.GET("/logout", r.auth.GetLogout)
}

// NewRouter returns a new router object
func NewRouter(admin Admin, auth AuthHander) *Router {
	return &Router{
		admin: admin,
		auth:  auth,
	}
}
