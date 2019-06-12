package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ribbonwall/be_ribbonwall/config"
	"github.com/ribbonwall/be_ribbonwall/endpoints"
	"github.com/ribbonwall/be_ribbonwall/models"
	singletonConfig "github.com/ribbonwall/be_ribbonwall/singletons/config"
	singletonDatabase "github.com/ribbonwall/be_ribbonwall/singletons/db"
	"github.com/ribbonwall/common/auth"
	logger "github.com/ribbonwall/common/logging"
	"time"
)

var (
	AdminGroup = "admin"
	configSing config.Config
	dbClient   *gorm.DB
)

func init() {
	// Init config
	configSing = singletonConfig.Get()

	// Init auth
	auth.InitAuth(configSing.Auth)

	// Init logger
	logger.NewLogger(configSing.Env)

	// Init MySQL Database
	var err error
	dbClient, err = singletonDatabase.Get()
	if err != nil {
		logger.Panicf("Could not init MySQL client: %v", err)
	}
	// gorm debug logging
	dbClient.LogMode(true)
	// gorm migrate up schema. Will only create tables and columns. Will not remove or modify existing columns.
	dbClient.AutoMigrate(&models.CompetitionResults{}, &models.Competitor{})
}

func main() {
	// Init gin router and routes
	router := gin.Default()

	// Cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Init gin middleware
	router.Use(EndpointMiddleware())

	// Init sessions
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// Init gin routes
	router.GET("/", endpoints.GetIndex)
	router.GET("/login", endpoints.LoginHandler)
	router.GET("/logout", endpoints.LogoutHandler)
	router.GET("/callback", endpoints.CallbackHandler)

	// Auth admin required
	router.GET("/competition_results", auth.Auth0Groups(AdminGroup), endpoints.GetCompetitionResults)
	router.POST("/competition_results/upload", auth.Auth0Groups(AdminGroup), endpoints.UploadCompetitionResults)

	_ = router.Run(":8080")
}

// Generic endpoint middleware
func EndpointMiddleware() gin.HandlerFunc {

	// Init Endpoints service for injecting dependencies
	endpointService := endpoints.NewEndpointService(
		dbClient,
		configSing,
	)

	return func(c *gin.Context) {
		c.Set("services", endpointService) // Available via `services := context.Keys["services"].(*Endpoints)`
		c.Next()
	}
}
