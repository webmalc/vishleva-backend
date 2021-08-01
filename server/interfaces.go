package server

import "github.com/gin-gonic/gin"

// InfoLogger logs errors.
type InfoLogger interface {
	Infof(format string, args ...interface{})
}

// Router is the router interface
type Router interface {
	BindRoutes(e *gin.Engine)
}
