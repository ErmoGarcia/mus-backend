package auth

import (
	"net/http"

	"github.com/ErmoGarcia/mus-backend/models"
	"github.com/ErmoGarcia/mus-backend/services/users"
	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm-password" binding:"required,eqfield=Password"`
}

func Register(c *gin.Context) {

	// bind input
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: validation

	// add user to db
	userService := users.GetService()

	user := models.User{}

	user.Username = input.Username
	user.Password = input.Password

	err := userService.CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "validated!"})

}
