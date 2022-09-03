package admin

import (
	"github.com/qor/admin"
)

// ResourceInitializer initializes a resource.
type ResourceInitializer interface {
	initMenu(a *admin.Admin)
	init(a *admin.Admin)
}
