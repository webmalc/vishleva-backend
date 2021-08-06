package admin

import (
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/validations"
	"github.com/webmalc/vishleva-backend/models"
)

type userResource struct{}

func (u *userResource) initMenu(a *admin.Admin) {
	a.AddMenu(&admin.Menu{Name: "Admins", Priority: -10})
}

func (u *userResource) init(a *admin.Admin) {
	usr := a.AddResource(&models.User{}, &admin.Config{Name: "Admins"})
	usr.IndexAttrs("ID", "Email", "LastLogin")
	usr.Meta(&admin.Meta{
		Name: "Password",
		Type: "password",
		Setter: func(
			resource interface{},
			metaValue *resource.MetaValue,
			context *qor.Context,
		) {
			values := metaValue.Value.([]string)
			if len(values) > 0 {
				pwd := values[0]
				if pwd == "" {
					return
				}
				u := resource.(*models.User)
				err := u.SetPassword(pwd)
				if err != nil {
					context.DB.AddError( // nolint // unnecessary: errcheck
						validations.NewError(usr, "Password", err.Error()),
					)
				}
			}
		},
	})
}
