package server

import "net/http"

// InfoLogger logs errors.
type InfoLogger interface {
	Infof(format string, args ...interface{})
}

// Admin is the admin interface
type Admin interface {
	Mount() *http.ServeMux
	GetBasePath() string
}
