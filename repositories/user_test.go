package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/common/test"
	"github.com/webmalc/vishleva-backend/models"
)

func TestUserRepository_LoginAndReturnUser(t *testing.T) {
	conn := db.NewConnection()
	email := "test@test.com"
	pass := "password"
	defer conn.Close()
	models.Migrate(conn)
	var users []models.User
	result := conn.Find(&users)
	user := models.User{Email: email}
	repo := NewUserRepository(conn.DB)
	err := user.SetPassword(pass)

	assert.Equal(t, int64(0), result.RowsAffected)
	assert.Empty(t, users)
	assert.Nil(t, err)
	conn.Create(&user)
	assert.Nil(t, user.LastLogin)

	_, err = repo.LoginAndReturnUser("invalid@user.com", pass)
	assert.Equal(t, "user is not found", err.Error())

	_, err = repo.LoginAndReturnUser(email, "invalid_password")
	assert.Equal(t, "password is incorrect", err.Error())

	loginUser, err := repo.LoginAndReturnUser(email, pass)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, loginUser.ID)
}

func TestNewUserRepository(t *testing.T) {
	c := db.NewConnection()
	defer c.Close()
	r := NewUserRepository(c.DB)
	assert.Equal(t, r.db, c.DB)
}

// Setups the tests
func TestMain(m *testing.M) {
	test.Run(m)
}
