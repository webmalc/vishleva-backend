package admin

import (
	"fmt"

	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/services"
)

type imageResource struct {
	tag        *admin.Resource
	collection *admin.Resource
}

func (r *imageResource) initMenu(a *admin.Admin) {
	a.AddMenu(&admin.Menu{Name: "Images", Priority: 1})
	a.AddMenu(&admin.Menu{Name: "Collections", Priority: 2})
	a.AddMenu(&admin.Menu{Name: "Tags", Priority: 3})
}

func (r *imageResource) initCollection(a *admin.Admin) {
	r.initTags(a)
	collection := a.AddResource(&models.Collection{})
	collection.IndexAttrs("-Description", "-Summary")
	collection.Meta(&admin.Meta{
		Name:   "Tags",
		Config: &admin.SelectManyConfig{SelectMode: "bottom_sheet"},
	})
	collection.Meta(&admin.Meta{
		Name: "Image",
		FormattedValuer: func(record interface{}, context *qor.Context) interface{} {
			if r, ok := record.(*models.Collection); ok && r.ImageID != nil {
				return fmt.Sprintf("image #%d", *r.ImageID)
			}

			return "-"
		},
		Config: &admin.SelectOneConfig{
			SelectMode: "bottom_sheet",
			AllowBlank: true,
		},
	})
	collection.SearchAttrs("Name", "Description", "Summary")
	collection.Filter(&admin.Filter{
		Name:   "Tags",
		Config: &admin.SelectOneConfig{RemoteDataResource: r.tag},
	})
	collection.Filter(&admin.Filter{Name: "IsEnabled"})
	r.collection = collection
}

func (r *imageResource) initTags(a *admin.Admin) {
	tag := a.AddResource(&models.Tag{})
	tag.IndexAttrs("ID", "Name", "Collections")
	tag.NewAttrs("Name")
	tag.EditAttrs("Name")
	tag.SearchAttrs("Name")
	r.tag = tag
}

func (r *imageResource) batchTags(argument *admin.ActionArgument) error {
	tags, ok := argument.Argument.(*models.Image)
	if !ok {
		panic("admin: assertion error.")
	}
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
}

func (r *imageResource) init(a *admin.Admin) {
	r.initCollection(a)

	image := a.AddResource(&models.Image{}, &admin.Config{PageCount: 50})
	image.IndexAttrs("File", "Tags")
	image.Meta(&admin.Meta{
		Name:   "Tags",
		Config: &admin.SelectManyConfig{SelectMode: "bottom_sheet"},
	})
	image.SearchAttrs("Name", "Description")
	image.Filter(&admin.Filter{
		Name:   "Tags",
		Config: &admin.SelectOneConfig{RemoteDataResource: r.tag},
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
		Handler:  r.batchTags,
		Modes:    []string{"batch"},
	})
	image.UseTheme("grid")
}
