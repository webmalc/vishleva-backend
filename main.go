//+build !test

package main

import (
	"github.com/webmalc/vishleva-backend/admin"
	"github.com/webmalc/vishleva-backend/admin/bindatafs"
	"github.com/webmalc/vishleva-backend/cmd"
	"github.com/webmalc/vishleva-backend/common/config"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/common/logger"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/server"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	conn := db.NewConnection()
	models.Migrate(conn)
	httpServer := server.NewServer(admin.NewAdmin(conn.DB), log)
	defer conn.Close()
	cmdRouter := cmd.NewCommandRouter(
		log,
		httpServer,
		bindatafs.NewGenerator(),
	)
	cmdRouter.Run()
}
