package controllers

import (
	"net/http"

	"github.com/LiveSchedule-v2/initializers"
	"github.com/LiveSchedule-v2/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginDiscord(c *gin.Context) {
	// Get the req
	var body struct {
		Email     string
		Password  string
		DiscordId string `form:"discord_id" json:"discord_id"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Look up requested user
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
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid password",
		})
		return
	}

	// Save discordId
	initializers.DB.Model(&user).Update("discord_id", body.DiscordId)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully register discord_id",
	})
}
