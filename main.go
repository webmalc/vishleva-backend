//go:build !test
// +build !test

package main

import (
	"github.com/webmalc/vishleva-backend/admin"
	"github.com/webmalc/vishleva-backend/admin/bindatafs"
	"github.com/webmalc/vishleva-backend/calendar"
	"github.com/webmalc/vishleva-backend/cmd"
	"github.com/webmalc/vishleva-backend/common/config"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/common/logger"
	"github.com/webmalc/vishleva-backend/common/session"
	"github.com/webmalc/vishleva-backend/handlers"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/repositories"
	"github.com/webmalc/vishleva-backend/routes"
	"github.com/webmalc/vishleva-backend/server"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	conn := db.NewConnection()
	sessionConfig := session.NewSession()
	userRepository := repositories.NewUserRepository(conn.DB)
	tariffRepository := repositories.NewTariffRepository(conn.DB)
	tagsRepository := repositories.NewTagRepository(conn.DB)
	reviewsRepository := repositories.NewReviewRepository(conn.DB)
	collectionsRepository := repositories.NewCollectionRepository(conn.DB)
	imagesRepository := repositories.NewImageRepository(conn.DB)
	orderRepository := repositories.NewOrderRepository(conn.DB)
	cal := calendar.NewGenerator(orderRepository)
	models.Migrate(conn)
	router := routes.NewRouter(
		admin.NewAdmin(conn.DB, sessionConfig),
		handlers.NewAuthHandler(sessionConfig, userRepository, log),
		handlers.NewTariffsHandler(tariffRepository),
		handlers.NewTagsHandler(tagsRepository),
		handlers.NewReviewsHandler(reviewsRepository),
		handlers.NewCollectionHandler(collectionsRepository),
		handlers.NewImagesHandler(imagesRepository),
		handlers.NewCalendarHandler(cal),
	)
	httpServer := server.NewServer(router, log, sessionConfig)
	defer conn.Close()
	cmdRouter := cmd.NewCommandRouter(
		log,
		httpServer,
		bindatafs.NewGenerator(),
	)
	cmdRouter.Run()
}
