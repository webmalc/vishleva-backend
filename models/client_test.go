package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/common/db"
)

func TestClient_Validate(t *testing.T) {
	conn := db.NewConnection()
	client := &Client{
		Name: "test",
	}
	client.Validate(conn.DB)
	assert.Len(t, conn.GetErrors(), 0)

	client.Name = ""
	client.Phone = "89251110022"
	client.Validate(conn.DB)
	assert.Len(t, conn.GetErrors(), 2)
	assert.Contains(t, conn.GetErrors()[0].Error(), "name is empty")
	assert.Contains(t, conn.GetErrors()[1].Error(), "start with 7")
}
