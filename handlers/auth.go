package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/webmalc/vishleva-backend/common/session"
)

// AuthHander is auth handler
type AuthHander struct {
	session     *session.Session
	userLoginer UserLoginer
	logger      ErrorLogger
	config      *Config
}

// GetLogin returns the login handler function
func (h *AuthHander) GetLogin(c *gin.Context) {
	s := sessions.Default(c)
	if sessions.Default(c).Get(h.session.Key) != nil {
		c.Redirect(http.StatusSeeOther, h.config.AdminPath)
		return
	}
	flashes := s.Flashes()
	if err := s.Save(); err != nil {
		h.logger.Errorf("Unable to save session: %v", err)
	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"flashes": flashes,
	})
}

// PostLogin is the handler to check if the user can connect
func (h *AuthHander) PostLogin(c *gin.Context) {
	s := sessions.Default(c)
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := h.userLoginer.LoginAndReturnUser(email, password)
	if err != nil {
		s.AddFlash("invalid login or password")
		if err := s.Save(); err != nil {
			h.logger.Errorf("Unable to save session: %v", err)
		}

		c.Redirect(http.StatusSeeOther, h.config.LoginPath)
		return
	}
	s.Set(h.session.Key, user.ID)
	err = s.Save()
	if err != nil {
		h.logger.Errorf("Unable to save session: %v", err)
		c.Redirect(http.StatusSeeOther, h.config.LoginPath)
		return
	}
	c.Redirect(http.StatusSeeOther, "/"+h.config.AdminPath)
}

// GetLogout allows the user to disconnect
func (h *AuthHander) GetLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Delete(h.session.Key)
	if err := s.Save(); err != nil {
		h.logger.Errorf("Unable to save session: %v", err)
	}
	c.Redirect(http.StatusSeeOther, h.config.LoginPath)
}

// NewAuthHandler returns a new router object
func NewAuthHandler(
	s *session.Session, u UserLoginer, l ErrorLogger,
) *AuthHander {
	return &AuthHander{
		session:     s,
		userLoginer: u,
		config:      NewConfig(),
		logger:      l,
	}
}
