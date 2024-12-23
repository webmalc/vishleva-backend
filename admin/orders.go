package admin

import (
	"github.com/qor/admin"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/utils"
)

type orderResource struct {
	config *Config
	client *admin.Resource
}

func (r *orderResource) initMenu(a *admin.Admin) {
	a.AddMenu(&admin.Menu{Name: "Orders", Priority: 1})
	a.AddMenu(&admin.Menu{Name: "Clients", Priority: 2}) //nolint // unnecessary: mnd
}

func (r *orderResource) initClient(a *admin.Admin) {
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
		Handler: utils.GetDateFilter("clients", "created_at"),
	})
	r.client = client
}

func (r *orderResource) init(a *admin.Admin) {
	r.initClient(a)

	order := a.AddResource(&models.Order{})
	order.IndexAttrs(
		"ID", "Name", "Client", "Begin", "End", "Total",
		"Paid", "Status", "Source", "CreatedAt", "UpdatedAt",
	)
	order.NewAttrs("-CreatedAt", "-UpdatedAt")
	order.EditAttrs("-CreatedAt", "-UpdatedAt")
	order.SearchAttrs(
		"Name", "Comment", "Client.Name",
		"Client.Email", "Client.Comment",
		"Client.Social", "Client.Phone",
	)
	order.Meta(&admin.Meta{Name: "Total", Type: "float"})
	order.Meta(&admin.Meta{Name: "Paid", Type: "float"})
	order.Meta(&admin.Meta{
		Name: "Status",
		Config: &admin.SelectOneConfig{
			Collection: r.config.OrderStatuses,
		}})
	order.Meta(&admin.Meta{
		Name: "Source",
		Config: &admin.SelectOneConfig{
			Collection: r.config.OrderSources,
		}})
	order.Meta(&admin.Meta{
		Name: "Client",
		Config: &admin.SelectOneConfig{
			SelectMode:         "bottom_sheet",
			AllowBlank:         true,
			RemoteDataResource: r.client,
		},
	})
	order.Filter(&admin.Filter{Name: "Begin"})
	order.Filter(&admin.Filter{Name: "End"})
	order.Filter(&admin.Filter{Name: "Status"})
	order.Filter(&admin.Filter{Name: "Source"})
	order.Filter(&admin.Filter{Name: "Client"})
	order.Filter(&admin.Filter{
		Name:    "CreatedAt",
		Handler: utils.GetDateFilter("orders", "created_at"),
	})
	order.Filter(&admin.Filter{
		Name:    "UpdatedAt",
		Handler: utils.GetDateFilter("orders", "updated_at"),
	})
}
