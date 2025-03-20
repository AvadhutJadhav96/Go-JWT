package helper

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/AvadhutJadhav96/Go-JWT/database"
	jwt "github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SignedDetails defines the structure of JWT claims that will be embedded in the token.
type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	User_type  string
	jwt.RegisteredClaims // Includes standard JWT claims such as expiration time
}

// MongoDB collection reference for the user table
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// Secret key used for signing JWTs, loaded from environment variables
var SECRET_KEY string = os.Getenv("SECRET_KEY")

// GenerateAllTokens creates both an access token and a refresh token for a user
func GenerateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error) {
	// Define the claims for the access token with an expiration time of 24 hours
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		User_type:  userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 24)), // Token expires in 24 hours
		},
	}

	// Define the claims for the refresh token with an expiration time of 7 days (168 hours)
	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 168)), // Token expires in 7 days
		},
	}

	// Generate the signed JWT tokens
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

// ValidateToken verifies the given JWT token and returns the claims if valid
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	// Parse the JWT token and extract claims
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	// Validate claims
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "The token is invalid"
		return
	}

	// Check if the token has expired
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now().Local()) {
		msg = "Token has expired"
		return
	}

	return claims, msg
}

// UpdateAllTokens updates the user's access and refresh tokens in the database
func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel() // Ensure the context is canceled to free resources

	var updateObj primitive.D

	// Update the token fields
	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refresh_token", signedRefreshToken})

	// Update the last modified timestamp
	Updated_at, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at", Updated_at})

	// Define the filter and update options
	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert, // Insert a new document if no match is found
	}

	// Perform the update operation
	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set", updateObj},
		},
		&opt,
	)

	if err != nil {
		log.Panic(err)
		return
	}
}
