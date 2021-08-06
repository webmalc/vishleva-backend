package admin

import (
	"github.com/qor/admin"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/services"
)

type eventResource struct{}

func (u *eventResource) initMenu(a *admin.Admin) {
	a.AddMenu(&admin.Menu{Name: "Clients", Priority: 1})
}

func (u *eventResource) init(a *admin.Admin) {
	client := a.AddResource(&models.Client{})
	client.IndexAttrs(
		"ID", "Name", "Phone", "Social",
		"Email", "Comment", "CreatedAt",
	)
	client.NewAttrs("-CreatedAt")
	client.EditAttrs("-CreatedAt")
	client.SearchAttrs(
		"Name", "Phone", "Social",
		"Email", "Comment",
	)
	client.Filter(&admin.Filter{
		Name:    "CreatedAt",
		Handler: services.GetDateFilter("clients", "created_at"),
	})
}
