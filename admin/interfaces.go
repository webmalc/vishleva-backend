package admin

import (
	"github.com/qor/admin"
)

// ResourceInitializer initializes a resource
type ResourceInitializer interface {
	init(a *admin.Admin)
}
