package routes

import (
	controller "github.com/AvadhutJadhav96/Go-JWT/controllers" // Import user-related controllers
	"github.com/AvadhutJadhav96/Go-JWT/middleware" // Import authentication middleware
	"github.com/gin-gonic/gin" // Import Gin framework for routing
)

// UserRoutes defines routes for fetching user data, protected by authentication middleware.
func UserRoutes(incomingRoutes *gin.Engine) {
	// Apply authentication middleware to all user-related routes
	incomingRoutes.Use(middleware.Authenticate())

	// Route to get a list of all users
	incomingRoutes.GET("/users", controller.GetUsers())

	// Route to get details of a specific user by user ID
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
