package admin

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/webmalc/vishleva-backend/admin/bindatafs"
	"github.com/webmalc/vishleva-backend/common/session"
)

// Admin is the admin structure
type Admin struct {
	config            *Config
	admin             *admin.Admin
	db                *gorm.DB
	session           *session.Session
	resourceFunctions []ResourceInitializer
}

// Init initializes the admin
func (a *Admin) Init() {
	a.admin = admin.New(&admin.AdminConfig{
		DB:       a.db,
		Auth:     newAuth(a.db, a.session),
		SiteName: a.config.SiteName,
	})
	a.admin.SetAssetFS(bindatafs.AssetFS.NameSpace("admin"))
	for _, resource := range a.resourceFunctions {
		resource.initMenu(a.admin)
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
func NewAdmin(db *gorm.DB, s *session.Session) *Admin {
	config := NewConfig()
	a := Admin{
		config:  config,
		db:      db,
		session: s,
		resourceFunctions: []ResourceInitializer{
			&reviewResource{},
			&tariffResource{},
			&userResource{},
			&imageResource{},
		},
	}
	a.Init()
	return &a
}
