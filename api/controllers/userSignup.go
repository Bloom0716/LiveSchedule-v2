package controllers

import (
	"net/http"

	"github.com/LiveSchedule-v2/initializers"
	"github.com/LiveSchedule-v2/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get the email, name, pass off req body
	var body struct {
		Email    string `form:"email"`
		Name     string `form:"name"`
		Password string `form:"password"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Check body
	if body.Email == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Required email and password",
		})
		return
	}

	//Hash the pass
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	//Create the user
	user := models.User{Email: body.Email, Name: body.Name, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Create user",
	})
}
