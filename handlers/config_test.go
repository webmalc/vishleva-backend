package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bitly/go-simplejson"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/webmalc/vishleva-backend/admin"
	"github.com/webmalc/vishleva-backend/calendar"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/common/logger"
	"github.com/webmalc/vishleva-backend/common/session"
	"github.com/webmalc/vishleva-backend/common/test"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/repositories"
	"github.com/webmalc/vishleva-backend/routes"
)

func createImages(conn *db.Database) {
	conn.Create(&models.Image{Name: "image one"})
	conn.Create(&models.Image{Name: "image two"})
}

func createCollections(conn *db.Database) {
	conn.Create(&models.Collection{Name: "collection one", IsEnabled: true})
	conn.Create(&models.Collection{Name: "collection two", IsEnabled: false})
}

func createTags(conn *db.Database) {
	conn.Create(&models.Tag{Name: "one"})
	conn.Create(&models.Tag{Name: "two"})
}

func createReviews(conn *db.Database) {
	conn.Create(&models.Review{Content: "one", IsEnabled: true})
	conn.Create(&models.Review{Content: "two", IsEnabled: false})
}

func createTariffs(conn *db.Database) {
	conn.Create(
		&models.Tariff{
			Name:         "one",
			Price:        decimal.NewFromInt(100), // nolint // unnecessary: unparam
			Duration:     60,
			Photos:       20,
			Retouch:      10,
			RetouchPrice: decimal.NewFromInt(10), // nolint // unnecessary: unparam
			IsEnabled:    true,
		},
	)
	conn.Create(
		&models.Tariff{
			Name:         "two",
			Price:        decimal.NewFromInt(33), // nolint // unnecessary: unparam
			Duration:     30,
			Photos:       10,
			Retouch:      10,
			RetouchPrice: decimal.NewFromInt(17), // nolint // unnecessary: unparam
			IsEnabled:    false,
		},
	)
}

func checkResponse(t *testing.T, url string, count int) {
	w, engine := initRoutes()
	req, _ := http.NewRequest("GET", url, nil)
	engine.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	json, err := simplejson.NewFromReader(w.Body)
	assert.Nil(t, err)
	entries, err := json.Get("entries").Array()
	assert.Nil(t, err)
	assert.Len(t, entries, count)
}

// initRoutes initializes the router
func initRoutes() (*httptest.ResponseRecorder, *gin.Engine) {
	log := logger.NewLogger()
	conn := db.NewConnection()
	sessionConfig := session.NewSession()
	userRepository := repositories.NewUserRepository(conn.DB)
	tariffsRepository := repositories.NewTariffRepository(conn.DB)
	tagsRepository := repositories.NewTagRepository(conn.DB)
	reviewsRepository := repositories.NewReviewRepository(conn.DB)
	collectionsRepository := repositories.NewCollectionRepository(conn.DB)
	imagesRepository := repositories.NewImageRepository(conn.DB)
	orderRepository := repositories.NewOrderRepository(conn.DB)
	cal := calendar.NewGenerator(orderRepository)
	models.Migrate(conn)
	router := routes.NewRouter(
		admin.NewAdmin(conn.DB, sessionConfig),
		NewAuthHandler(sessionConfig, userRepository, log),
		NewTariffsHandler(tariffsRepository),
		NewTagsHandler(tagsRepository),
		NewReviewsHandler(reviewsRepository),
		NewCollectionHandler(collectionsRepository),
		NewImagesHandler(imagesRepository),
		NewCalendarHandler(cal),
	)

	engine := gin.Default()
	engine.LoadHTMLGlob("app/views/auth/*")
	engine.Use(sessions.Sessions(sessionConfig.Name, sessionConfig.Store))
	router.BindRoutes(engine)
	w := httptest.NewRecorder()
	createTariffs(conn)
	createTags(conn)
	createReviews(conn)
	createCollections(conn)
	createImages(conn)
	return w, engine
}

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, "admin", c.AdminPath)
	assert.Equal(t, "/auth/login", c.LoginPath)
}

// Setups the tests
func TestMain(m *testing.M) {
	if err := os.Chdir("../"); err != nil {
		panic(err)
	}
	test.Run(m)
}
