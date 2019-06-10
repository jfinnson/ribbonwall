package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/ribbonwall/be_ribbonwall/endpoints"
	"github.com/ribbonwall/be_ribbonwall/models"
	singletonConfig "github.com/ribbonwall/be_ribbonwall/singletons/config"
	singletonDatabase "github.com/ribbonwall/be_ribbonwall/singletons/db"
	logger "github.com/ribbonwall/common/logging"
)

var (
	listenAddr string
	healthy    int32
)

func main() {
	flag.StringVar(&listenAddr, "listen-addr", ":8080", "server listen address")
	flag.Parse()

	// Init gin router and routes
	router := gin.Default()

	// Init gin middleware
	router.Use(EndpointMiddleware())

	// Init gin routes
	router.GET("/", endpoints.GetIndex)
	router.GET("/healthz", healthz)
	router.GET("/competition_results", endpoints.GetCompetitionResults)
	router.POST("/competition_results/upload", endpoints.UploadCompetitionResults)

	_ = router.Run(":8080")
}

// Basic health endpoint
func healthz(c *gin.Context) {
	if atomic.LoadInt32(&healthy) == 1 {
		c.Status(http.StatusNoContent)
		return
	}
	c.Status(http.StatusServiceUnavailable)
}

// Generic endpoint middleware
func EndpointMiddleware() gin.HandlerFunc {
	// Get config settings
	config := singletonConfig.Get()

	// Init logger
	logger.NewLogger(config.Env)

	// Init MySQL Database
	dbClient, err := singletonDatabase.Get()
	if err != nil {
		log.Panicf("Could not init MySQL client: %v", err)
	}
	// gorm debug logging
	dbClient.LogMode(true)
	// gorm migrate up schema. Will only create tables and columns. Will not remove or modify existing columns.
	dbClient.AutoMigrate(&models.CompetitionResults{}, &models.Competitor{})

	// Init Endpoints service for injecting dependencies
	endpointService := endpoints.NewEndpointService(
		dbClient,
		config,
	)

	return func(c *gin.Context) {
		c.Set("services", endpointService) // Available via `services := context.Keys["services"].(*Endpoints)`
		c.Next()
	}
}
