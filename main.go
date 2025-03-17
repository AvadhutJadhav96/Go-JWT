package main 

import (
	"os"

	"github.com/gin-gonic/gin"   // Importing Gin framework
	"github.com/AvadhutJadhav96/Go-JWT/routes"              // Importing the routes package
)

func main() {
	// Get the PORT environment variable, or default to "8000"
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Create a new Gin router instance
	router := gin.New()

	// Use the Gin Logger middleware to log incoming HTTP requests
	router.Use(gin.Logger())

	// Register authentication routes (e.g., login, register)
	routes.AuthRoutes(router)

	// Register user-related routes (e.g., get user, update user)
	routes.UserRoutes(router)

	// Define a simple GET endpoint at /api-1
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	// Define another GET endpoint at /api-2
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	// Start the server and listen on the specified port
	router.Run(":" + port) // Corrected: `Run` (not `RUN`)
}
