package main

import (
	"github.com/gin-gonic/gin"
	"go-web/config"
	"go-web/router"
	"log"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Load configuration
	config.LoadConfig()

	// Get port from environment variable
	port := config.GetEnv("PORT", "8080")

	// Initialize routes
	router.InitRoutes(r)

	// Start the server
	log.Printf("server running on port %s...", port)
	r.Run(":" + port)
}
