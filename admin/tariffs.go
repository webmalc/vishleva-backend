package admin

import (
	"github.com/qor/admin"
	"github.com/webmalc/vishleva-backend/models"
)

type tariffResource struct{}

func (u *tariffResource) initMenu(a *admin.Admin) {
	a.AddMenu(&admin.Menu{Name: "Tariffs", Priority: 1})
}

func (u *tariffResource) init(a *admin.Admin) {
	tariff := a.AddResource(&models.Tariff{})
	tariff.IndexAttrs("-Description")
	tariff.Meta(&admin.Meta{Name: "Price", Type: "float"})
	tariff.Meta(&admin.Meta{
		Name:  "RetouchPrice",
		Type:  "float",
		Label: "Retouch price per piece",
	})
	tariff.Meta(&admin.Meta{Name: "Duration", Label: "Duration in minutes"})
	tariff.Meta(&admin.Meta{Name: "Photos", Label: "Number of photos"})
	tariff.Meta(&admin.Meta{
		Name: "Retouch", Label: "Number of retouched photos",
	})
}
