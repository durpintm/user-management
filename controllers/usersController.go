package controllers

import (
	"net/http"

	"github.com/durpintm/user-management/initializers"
	"github.com/durpintm/user-management/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var body struct {
	Email    string
	Password string
}

func SignUp(c *gin.Context) {

	// Get the email/password from the request body
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read the body"})
		return
	}

	// Hash the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create the user
	user := models.User{Email: body.Email, PasswordHash: string(hashPassword)}
	result := initializers.DB.Create(&user) // pass pointer of data to create user

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}

// Login
