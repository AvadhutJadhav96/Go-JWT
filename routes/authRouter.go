package routes

import (
	controller "github.com/AvadhutJadhav96/Go-JWT/controllers" // Import user authentication controllers
	"github.com/gin-gonic/gin" // Import Gin framework for routing
)

// AuthRoutes defines authentication-related routes for user signup and login.
func AuthRoutes(incomingRoutes *gin.Engine) {
	// Route for user signup
	incomingRoutes.POST("users/signup", controller.Signup())

	// Route for user login
	incomingRoutes.POST("users/login", controller.Login())
}
