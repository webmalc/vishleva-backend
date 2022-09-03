package admin

import (
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/validations"
	"github.com/webmalc/vishleva-backend/models"
)

type userResource struct {
	user *admin.Resource
}

func (r *userResource) initMenu(a *admin.Admin) {
	a.AddMenu(&admin.Menu{Name: "Admins", Priority: -10})
	a.AddMenu(&admin.Menu{
		Name:     "Clear cache",
		Link:     "/api/cache",
		Priority: -20,
	})
}

func (r *userResource) passwordSetter(
	res interface{},
	metaValue *resource.MetaValue,
	context *qor.Context,
) {
	values, ok := metaValue.Value.([]string)
	if !ok {
		panic("admin: assertion error.")
	}
	if len(values) > 0 {
		pwd := values[0]
		if pwd == "" {
			return
		}
		u, ok := res.(*models.User)
		if !ok {
			panic("admin: assertion error.")
		}
		err := u.SetPassword(pwd)
		if err != nil {
			context.DB.AddError( // nolint // unnecessary: errcheck
				validations.NewError(r.user, "Password", err.Error()),
			)
		}
	}
}

func (r *userResource) init(a *admin.Admin) {
	r.user = a.AddResource(&models.User{}, &admin.Config{Name: "Admins"})
	r.user.IndexAttrs("ID", "Email", "LastLogin")
	r.user.Meta(&admin.Meta{
		Name:   "Password",
		Type:   "password",
		Setter: r.passwordSetter,
	})
}
