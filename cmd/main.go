package main

import (
	"log"
	"svnfluence/internal/api"
	"svnfluence/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration (e.g., OpenAI API key)
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Set up Gin router
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	// Initialize API handlers with config
	api.SetupRoutes(router, cfg)

	// Start the server
	log.Println("Server starting on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
