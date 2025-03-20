package helper

import (
	"errors"
	"github.com/gin-gonic/gin" // Import Gin framework for handling request context
)

// CheckUserType verifies if the requesting user's role matches the required role.
func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type") // Retrieve the user type from context
	err = nil

	// If the user type does not match the required role, return an error
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}

// MatchUserTypeToUid ensures that a user with type "USER" can only access their own data.
func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type") // Retrieve the user type from context
	uid := c.GetString("uid") // Retrieve the user ID from context
	err = nil

	// If the user is of type "USER" but their ID does not match the provided ID, deny access
	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	// Check if the user has the required user type
	err = CheckUserType(c, userType)
	return err
}
