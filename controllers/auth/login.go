package auth

import (
	"net/http"

	"github.com/ErmoGarcia/mus-backend/services/users"
	"github.com/ErmoGarcia/mus-backend/utils/token"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	// bind input
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: validation

	// get user from db
	userService := users.GetService()

	user, err := userService.GetUserByName(input.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user doesn't exist."})
		return
	}

	// check password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	// generate auth token
	token, err := token.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func CurrentUser(c *gin.Context) {

	// get uid from token
	user_id, err := token.GetTokenID(token.ExtractToken(c))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get user data
	userService := users.GetService()
	user, err := userService.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}
