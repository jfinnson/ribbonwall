package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ribbonwall/common/auth"
	logger "github.com/ribbonwall/common/logging"
	"github.com/ribbonwall/domains/be_ribbonwall/config"
	"github.com/ribbonwall/domains/be_ribbonwall/endpoints"
	"github.com/ribbonwall/domains/be_ribbonwall/models"
	singletonConfig "github.com/ribbonwall/domains/be_ribbonwall/singletons/config"
	singletonDatabase "github.com/ribbonwall/domains/be_ribbonwall/singletons/db"
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

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("domains/fe_competitors/build", true)))

	// Init gin routes
	api := router.Group("/api/v1")
	{
		api.GET("/", endpoints.GetPing)
		api.GET("/login", endpoints.LoginHandler)
		api.GET("/logout", endpoints.LogoutHandler)
		api.GET("/callback", endpoints.CallbackHandler)
	}

	// Auth admin required
	adminApi := router.Group("/api_admin/v1")
	{
		// CRUD competitors
		api.GET("/competitors", endpoints.GetCompetitors)
		api.POST("/competitors", endpoints.CreateCompetitor)
		api.PUT("/competitors/:uuid", endpoints.UpdateCompetitor)
		api.DELETE("/competitors/:uuid", endpoints.DeleteCompetitor)

		// CRUD competition results
		api.GET("/competition_results", endpoints.GetCompetitionResults)
		api.POST("/competition_results", endpoints.CreateCompetitionResult)
		//api.PUT("/competition_results/:uuid", endpoints.UpdateCompetitionResult)
		//api.DELETE("/competition_results/:uuid", endpoints.DeleteCompetitionResult)

		adminApi.POST("/competition_results/upload", auth.Auth0Groups(AdminGroup), endpoints.UploadCompetitionResults)
	}

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
