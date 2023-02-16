package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Admin is the admin interface.
type Admin interface {
	Mount() *http.ServeMux
	GetBasePath() string
}

// AuthHandler is the handler interface.
type AuthHandler interface {
	GetLogin(c *gin.Context)
	PostLogin(c *gin.Context)
	GetLogout(c *gin.Context)
}

// ListHandler is the handler interface.
type ListHandler interface {
	GetList(c *gin.Context)
}

// PostHandler is the handler interface.
type PostHandler interface {
	Post(c *gin.Context)
}
