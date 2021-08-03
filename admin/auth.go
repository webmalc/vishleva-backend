package admin

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/webmalc/vishleva-backend/common/session"
	"github.com/webmalc/vishleva-backend/models"
)

// The admin auth
type auth struct {
	db      *gorm.DB
	session *session.Session
	config  *Config
}

// LoginURL satisfies the Auth interface and returns the route used to log
func (a auth) LoginURL(c *admin.Context) string { // nolint // unnecessary: unparam
	return a.config.LoginPath
}

// LogoutURL satisfies the Auth interface and returns the route used to logout
func (a auth) LogoutURL(c *admin.Context) string { // nolint // unnecessary: unparam
	return a.config.LogoutPath
}

// GetCurrentUser satisfies the Auth interface and returns the current user
func (a auth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	var userid uint

	s, err := a.session.Store.Get(c.Request, a.session.Name)
	if err != nil {
		return nil
	}
	if v, ok := s.Values[a.session.Key]; ok {
		userid = v.(uint)
	} else {
		return nil
	}

	var user models.User
	if !a.db.First(&user, "id = ?", userid).RecordNotFound() {
		return &user
	}
	return nil
}

// newAuth creates a new auth structure
func newAuth(db *gorm.DB, s *session.Session) *auth {
	return &auth{db: db, session: s, config: NewConfig()}
}
