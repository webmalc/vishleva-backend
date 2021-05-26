package admin

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/webmalc/vishleva-backend/admin/bindatafs"
)

// Admin is the admin structure
type Admin struct {
	config            *Config
	admin             *admin.Admin
	db                *gorm.DB
	resourceFunctions []ResourceInitializer
}

// Init initializes the admin
func (a *Admin) Init() {
	a.admin = admin.New(&admin.AdminConfig{DB: a.db})
	a.admin.SetAssetFS(bindatafs.AssetFS.NameSpace("admin"))
	for _, resource := range a.resourceFunctions {
		resource.init(a.admin)
	}
}

// GetBasePath returns the base admin path
func (a *Admin) GetBasePath() string {
	return a.config.AdminPath
}

// Mount mounts the admin
func (a *Admin) Mount() *http.ServeMux {
	mux := http.NewServeMux()
	a.admin.MountTo(a.GetBasePath(), mux)
	return mux
}

// NewAdmin returns a new admin object
func NewAdmin(db *gorm.DB) *Admin {
	config := NewConfig()
	a := Admin{
		config: config,
		db:     db,
		resourceFunctions: []ResourceInitializer{
			&userResource{},
		},
	}
	a.Init()
	return &a
}
