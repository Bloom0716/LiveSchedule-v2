package controllers

import (
	"net/http"

	"github.com/LiveSchedule-v2/initializers"
	"github.com/LiveSchedule-v2/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginDiscordUser(c *gin.Context) {
	// Get the req
	var body struct {
		Email     string
		Password  string
		DiscordId int `json:"discord_id"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Look up user
	user := models.User{}
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email",
		})
		return
	}

	// Compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}

	// Create DiscordUser
	discordUser := models.DiscordUser{
		UserId:    user.ID,
		DiscordId: body.DiscordId,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
	}
	result := initializers.DB.Create(&discordUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create DiscordUser",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Create DiscordUser",
	})
}
