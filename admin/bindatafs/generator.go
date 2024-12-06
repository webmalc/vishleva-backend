package bindatafs

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// Generator is the bindatafs structure.
type Generator struct{}

// Run runs the generator.
func (g *Generator) Run(_ []string) {
	assetFS := AssetFS.NameSpace("admin")

	// Register view paths into AssetFS
	path := os.Getenv("GOPATH") + "/src/github.com/qor/admin/views"
	if err := assetFS.RegisterPath(path); err != nil {
		panic(errors.Wrap(err, "admin"))
	}
	fmt.Printf("The path has been added: %s\n", path)

	// Compile templates under registered view paths into binary
	if err := assetFS.Compile(); err != nil {
		panic(errors.Wrap(err, "admin"))
	}
}

// NewGenerator return a new generator.
func NewGenerator() *Generator {
	return &Generator{}
}
