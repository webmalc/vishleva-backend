package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Admin is the admin interface
type Admin interface {
	Mount() *http.ServeMux
	GetBasePath() string
}

// AuthHander is the handler interface
type AuthHander interface {
	GetLogin(c *gin.Context)
	PostLogin(c *gin.Context)
	GetLogout(c *gin.Context)
}
