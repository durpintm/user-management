package controllers

import (
	"crypto/rand"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/durpintm/user-management/initializers"
	"github.com/durpintm/user-management/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var body struct {
	Email          string
	Password       string
	InvitationCode string
}

// Signup
func SignUp(c *gin.Context) {

	// Get the email/password from the request body
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read the body"})
		return
	}

	// Validate the invitation code
	var inviteCode models.InvitationCode
	initializers.DB.First(&inviteCode, "invitation_code = ?", body.InvitationCode)

	// If no models.InvitationCode record or IsUsed is true => validation failed
	if inviteCode.ID == 0 || inviteCode.IsUsed {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not validate the invitation code"})
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

	// Set IsUsed to true and delete the InvitationCode record
	initializers.DB.Model(&inviteCode).Updates(models.InvitationCode{IsUsed: true})

	// UPDATE InvitationCode SET deleted_at={current_time} WHERE ID = inviteCode.ID; => Soft Delete
	initializers.DB.Delete(&inviteCode)

	// Respond
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}

// Login
<<<<<<< HEAD

// GenerateRandomToken generates a random code of length n
func generateRandomCode(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, length)
	for i := range token {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		token[i] = charset[idx.Int64()]
	}
	return string(token), nil
=======
func Login(c *gin.Context) {

	// Get the email and password from the request body
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read the request body"})
		return
	}

	// Look up for registered user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare sent in password with the registered user password hash
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid email or password"})
		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", true, true)
	// Send back the JWT token
	c.JSON(http.StatusOK, gin.H{})
>>>>>>> 9da78b71a424334ce391e1f41bd89e908393c453
}
