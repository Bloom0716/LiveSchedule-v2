package controllers

import (
	"net/http"

	"github.com/LiveSchedule-v2/initializers"
	"github.com/LiveSchedule-v2/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UpdateUser(c *gin.Context) {
	// Get the req
	userId, err := c.Cookie("UserId")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	DB := initializers.DB
	var body struct {
		Name     string `form:"name"`
		Password string `form:"password"`
	}

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read userId",
		})
		return
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Look up requested user
	user := models.User{}
	DB.First(&user, "ID = ?", userId)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	// Validation request
	if body.Name == "" && body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Require name or password",
		})
		return
	}

	// Update
	if body.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to hash password",
			})
			return
		}

		DB.Model(&user).Updates(models.User{Name: body.Name, Password: string(hash)})
	} else {
		DB.Model(&user).Update("name", body.Name)
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully update user data",
		"data": gin.H{
			"userId": user.ID,
			"name":   user.Name,
			"email":  user.Email,
		},
	})
}
