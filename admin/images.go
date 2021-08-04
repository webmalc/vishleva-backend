package admin

import (
	"github.com/qor/admin"
	"github.com/webmalc/vishleva-backend/models"
)

type imageResource struct{}

func (u *imageResource) init(a *admin.Admin) {
	tag := a.AddResource(&models.Tag{})
	tag.IndexAttrs("ID", "Name", "Collections")
	tag.NewAttrs("Name")
	tag.EditAttrs("Name")

	collection := a.AddResource(&models.Collection{})

	collection.Filter(&admin.Filter{
		Name:   "Tags",
		Config: &admin.SelectOneConfig{RemoteDataResource: tag},
	})
	collection.Filter(&admin.Filter{Name: "IsEnabled"})

	collection.Meta(&admin.Meta{
		Name:   "Tags",
		Config: &admin.SelectManyConfig{SelectMode: "bottom_sheet"},
	})
}
