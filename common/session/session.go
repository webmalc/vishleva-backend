package session

import (
	"github.com/gin-contrib/sessions/cookie"
)

// Session is the session config
type Session struct {
	Name  string
	Key   string
	Store cookie.Store
}

// NewSession provides a new session config object
func NewSession() *Session {
	config := NewConfig()
	return &Session{Name: config.SessionName,
		Key:   config.SessionKey,
		Store: cookie.NewStore([]byte(config.SessionSecret)),
	}
}
