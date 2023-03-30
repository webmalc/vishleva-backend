//go:build !test
// +build !test

package main

import (
	"github.com/webmalc/vishleva-backend/common/config"
	"github.com/webmalc/vishleva-backend/common/logger"
	"github.com/webmalc/vishleva-backend/common/messenger"
)

type MessageGetter struct{}

func (m *MessageGetter) GetText() string {
	return "text"
}
func (m *MessageGetter) GetHTML() string {
	return "html"
}
func (m *MessageGetter) GetSubject() string {
	return "subject"
}

type ContactsGetter struct{}

func (c *ContactsGetter) GetEmail() string {
	return "email"
}
func (c *ContactsGetter) GetTelegram() string {
	return "telegram"
}
func (c *ContactsGetter) GetVK() string {
	return "vk"
}

func main() {

	config.Setup()
	log := logger.NewLogger()
	m := messenger.NewMessenger(log)
	c := &ContactsGetter{}
	mg := &MessageGetter{}
	m.Send(c, mg)

	// config.Setup()
	// log := logger.NewLogger()
	// conn := db.NewConnection()
	// sessionConfig := session.NewSession()
	// userRepository := repositories.NewUserRepository(conn.DB)
	// tariffRepository := repositories.NewTariffRepository(conn.DB)
	// tagsRepository := repositories.NewTagRepository(conn.DB)
	// reviewsRepository := repositories.NewReviewRepository(conn.DB)
	// collectionsRepository := repositories.NewCollectionRepository(conn.DB)
	// imagesRepository := repositories.NewImageRepository(conn.DB)
	// orderRepository := repositories.NewOrderRepository(conn.DB)
	// cal := calendar.NewGenerator(orderRepository)
	// models.Migrate(conn)
	// router := routes.NewRouter(
	// 	admin.NewAdmin(conn.DB, sessionConfig),
	// 	handlers.NewAuthHandler(sessionConfig, userRepository, log),
	// 	handlers.NewTariffsHandler(tariffRepository),
	// 	handlers.NewTagsHandler(tagsRepository),
	// 	handlers.NewReviewsHandler(reviewsRepository),
	// 	handlers.NewCollectionHandler(collectionsRepository),
	// 	handlers.NewImagesHandler(imagesRepository),
	// 	handlers.NewCalendarHandler(cal),
	// 	handlers.NewBookHandler(
	// 		services.NewBookingService(
	// 			log,
	// 			repositories.NewClientRepository(conn.DB),
	// 			orderRepository,
	// 		),
	// 	),
	// )
	// httpServer := server.NewServer(router, log, sessionConfig)
	// defer conn.Close()
	// cmdRouter := cmd.NewCommandRouter(
	// 	log,
	// 	httpServer,
	// 	bindatafs.NewGenerator(),

	// )
	// cmdRouter.Run()
}
