package routes

import (
	"net/http"
)

// Admin is the admin interface
type Admin interface {
	Mount() *http.ServeMux
	GetBasePath() string
}
