package admin

import (
	"net/http"
	"testing"

	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/common/session"
)

func Test_auth_LoginURL(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	ses := session.NewSession()
	a := newAuth(conn.DB, ses)
	assert.Equal(t, "/auth/login", a.LoginURL(&admin.Context{}))
}

func Test_auth_LogoutURL(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	ses := session.NewSession()
	a := newAuth(conn.DB, ses)
	assert.Equal(t, "/auth/logout", a.LogoutURL(&admin.Context{}))
}

func Test_auth_GetCurrentUser(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	ses := session.NewSession()
	context := &admin.Context{
		Context: &qor.Context{Request: &http.Request{}},
	}
	a := newAuth(conn.DB, ses)
	assert.Nil(t, a.GetCurrentUser(context))
}

func Test_newAuth(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	ses := session.NewSession()
	a := newAuth(conn.DB, ses)
	assert.NotNil(t, a.db)
	assert.NotNil(t, a.session)
}
