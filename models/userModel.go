package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive" // Import MongoDB's BSON ObjectID package
)

// User struct represents the schema for the user collection in MongoDB.
type User struct {
	ID            primitive.ObjectID `bson:"_id"` // Unique identifier for the user, stored as ObjectID in MongoDB
	First_name    *string            `json:"first_name" validate:"required,min=2,max=100"` // User's first name (required, min 2, max 100 characters)
	Last_name     *string            `json:"last_name" validate:"required,min=2,max=100"` // User's last name (required, min 2, max 100 characters)
	Password      *string            `json:"Password" validate:"required,min=6"` // User's password (required, minimum 6 characters)
	Email         *string            `json:"email" validate:"email,required"` // User's email (required, must be a valid email format)
	Phone         *string            `json:"phone" validate:"required"` // User's phone number (required)
	Token         *string            `json:"token"` // JWT token for authentication
	User_type     *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"` // User type (must be either ADMIN or USER)
	Refresh_token *string            `json:"refresh_token"` // Refresh token for JWT authentication
	Created_at    time.Time          `json:"created_at"` // Timestamp when the user was created
	Updated_at    time.Time          `json:"updated_at"` // Timestamp when the user details were last updated
	User_id       string             `json:"user_id"` // Unique user ID (stored as a string)
}
