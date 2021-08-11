package handlers

import (
	"net/http/httptest"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
	"github.com/webmalc/vishleva-backend/admin"
	"github.com/webmalc/vishleva-backend/common/db"
	"github.com/webmalc/vishleva-backend/common/logger"
	"github.com/webmalc/vishleva-backend/common/session"
	"github.com/webmalc/vishleva-backend/models"
	"github.com/webmalc/vishleva-backend/repositories"
	"github.com/webmalc/vishleva-backend/routes"
)

// Config is the configuration struct.
type Config struct {
	AdminPath string
	LoginPath string
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

// initRoutes initializes the router
func initRoutes() (*httptest.ResponseRecorder, *gin.Engine) {
	log := logger.NewLogger()
	conn := db.NewConnection()
	sessionConfig := session.NewSession()
	userRepository := repositories.NewUserRepository(conn.DB)
	tariffsRepository := repositories.NewTariffRepository(conn.DB)
	models.Migrate(conn)
	router := routes.NewRouter(
		admin.NewAdmin(conn.DB, sessionConfig),
		NewAuthHandler(sessionConfig, userRepository, log),
		NewTariffsHandler(tariffsRepository),
	)

	engine := gin.Default()
	engine.LoadHTMLGlob("app/views/auth/*")
	engine.Use(sessions.Sessions(sessionConfig.Name, sessionConfig.Store))
	router.BindRoutes(engine)
	w := httptest.NewRecorder()
	createTariffs(conn)
	return w, engine
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	config := &Config{
		AdminPath: viper.GetString("admin_path"),
		LoginPath: viper.GetString("login_path"),
	}
	return config
}
