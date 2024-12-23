package admin

import (
	"fmt"

	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/webmalc/vishleva-backend/models"
)

type reviewResource struct{}

func (u *reviewResource) initMenu(a *admin.Admin) {
	a.AddMenu(&admin.Menu{Name: "Reviews", Priority: 1})
}

func (u *reviewResource) init(a *admin.Admin) {
	images := a.AddResource(&models.Image{}, &admin.Config{Invisible: true})
	images.IndexAttrs("ID", "File", "Name", "Tags")
	images.SearchAttrs("Name", "Description", "Tags")

	client := a.AddResource(&models.Client{}, &admin.Config{Invisible: true})
	client.IndexAttrs("ID", "Name", "Social", "Email", "Phone", "Comment")

	reviews := a.AddResource(&models.Review{})
	reviews.IndexAttrs("ID", "Content", "Client", "Image", "IsEnabled")
	reviews.Meta(&admin.Meta{
		Name: "Image",
		FormattedValuer: func(record interface{}, _ *qor.Context) interface{} {
			if r, ok := record.(*models.Review); ok && r.ImageID != nil {
				return fmt.Sprintf("image #%d", *r.ImageID)
			}

			return "-"
		},
		Config: &admin.SelectOneConfig{
			SelectMode:         "bottom_sheet",
			AllowBlank:         true,
			RemoteDataResource: images,
		},
	})
	reviews.Meta(&admin.Meta{
		Name: "Client",
		Config: &admin.SelectOneConfig{
			SelectMode:         "bottom_sheet",
			AllowBlank:         true,
			RemoteDataResource: client,
		},
	})
}
