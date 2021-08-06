package admin

import (
	"github.com/qor/admin"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/services"
)

type imageResource struct{}

func (u *imageResource) initMenu(a *admin.Admin) {
	a.AddMenu(&admin.Menu{Name: "Images", Priority: 10})
	a.AddMenu(&admin.Menu{Name: "Collections", Priority: 20})
	a.AddMenu(&admin.Menu{Name: "Tags", Priority: 30})
}

func (u *imageResource) init(a *admin.Admin) {
	// tags
	tag := a.AddResource(&models.Tag{})
	tag.IndexAttrs("ID", "Name", "Collections")
	tag.NewAttrs("Name")
	tag.EditAttrs("Name")
	tag.SearchAttrs("Name")

	collection := a.AddResource(&models.Collection{})
	collection.IndexAttrs("-Description", "-Summary")
	collection.Meta(&admin.Meta{
		Name:   "Tags",
		Config: &admin.SelectManyConfig{SelectMode: "bottom_sheet"},
	})
	collection.SearchAttrs("Name", "Description", "Summary")
	collection.Filter(&admin.Filter{
		Name:   "Tags",
		Config: &admin.SelectOneConfig{RemoteDataResource: tag},
	})
	collection.Filter(&admin.Filter{Name: "IsEnabled"})

	// images
	image := a.AddResource(&models.Image{}, &admin.Config{PageCount: 50})
	image.IndexAttrs("File", "Tags")
	image.Meta(&admin.Meta{
		Name:   "Tags",
		Config: &admin.SelectManyConfig{SelectMode: "bottom_sheet"},
	})
	image.SearchAttrs("Name", "Description")
	image.Filter(&admin.Filter{
		Name:   "Tags",
		Config: &admin.SelectOneConfig{RemoteDataResource: tag},
	})
	image.Filter(&admin.Filter{
		Name:    "CreatedAt",
		Handler: services.GetDateFilter("images", "created_at"),
	})
	image.Filter(&admin.Filter{
		Name:    "UpdatedAt",
		Handler: services.GetDateFilter("images", "updated_at"),
	})

	batchTags := *image
	batchTags.EditAttrs("Tags")
	image.Action(&admin.Action{
		Name:     "Update tags",
		Resource: &batchTags,
		Handler: func(argument *admin.ActionArgument) error {
			tags := argument.Argument.(*models.Image)
			for _, record := range argument.FindSelectedRecords() {
				if newImage, ok := record.(*models.Image); ok {
					if len(tags.Tags) > 0 {
						argument.Context.
							GetDB().
							Model(&newImage).
							Association("Tags").
							Replace(tags.Tags)
					}
				}
			}
			return nil
		},
		Modes: []string{"batch"},
	})
	image.UseTheme("grid")
}
