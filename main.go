package main

import (
	routes "github.com/AvadhutJadhav96/Go-JWT/routes" // Importing route handlers
	"os"
	"log"
	"github.com/gin-gonic/gin"       // Gin framework for handling HTTP requests
	"github.com/joho/godotenv"       // Package to load environment variables from a .env file
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file") // Log and exit if the .env file cannot be loaded
	}

	// Get the PORT value from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default to port 8000 if not specified in .env
	}

	// Create a new Gin router instance
	router := gin.New()
	router.Use(gin.Logger()) // Use Gin's built-in logging middleware

	// Register authentication and user routes
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// Define a test API endpoint - api-1
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	// Define another test API endpoint - api-2
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	// Start the server on the specified port
	router.Run(":" + port)
}
