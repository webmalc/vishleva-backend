package admin

import (
	"fmt"

	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/webmalc/vishleva-backend/models"
)

type reviewResource struct{}

func (u *reviewResource) initMenu(a *admin.Admin) {
	a.AddMenu(&admin.Menu{Name: "Reviews", Priority: 50})
}

func (u *reviewResource) init(a *admin.Admin) {
	images := a.AddResource(&models.Image{}, &admin.Config{Invisible: true})
	images.IndexAttrs("ID", "File", "Name", "Tags")
	images.SearchAttrs("Name", "Description", "Tags")

	reviews := a.AddResource(&models.Review{})
	reviews.IndexAttrs("ID", "Content", "Image", "IsEnabled")
	reviews.Meta(&admin.Meta{
		Name: "Image",
		FormattedValuer: func(record interface{}, context *qor.Context) interface{} {
			if r, ok := record.(*models.Review); ok && r.ImageID > 0 {
				return fmt.Sprintf("image #%d", r.ImageID)
			}
			return "-"
		},
		Config: &admin.SelectOneConfig{
			SelectMode:         "bottom_sheet",
			AllowBlank:         true,
			RemoteDataResource: images,
		},
	})
}
