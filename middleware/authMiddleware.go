package middleware

import (
	"fmt"
	"net/http"
	helper "github.com/AvadhutJadhav96/Go-JWT/helpers" // Import helper functions for token validation
	"github.com/gin-gonic/gin" // Import Gin framework for middleware handling
)

// Authenticate is a middleware function that verifies JWT tokens in incoming requests.
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the token from the request header
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			// If no token is provided, return an error response
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort() // Stop further request processing
			return
		}

		// Validate the provided token
		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			// If token validation fails, return an error response
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort() // Stop further request processing
			return
		}

		// Store the extracted claims in the request context for further use
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)

		// Proceed to the next middleware or handler
		c.Next()
	}
}
