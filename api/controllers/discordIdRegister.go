package controllers

import (
	"net/http"

	"github.com/LiveSchedule-v2/initializers"
	"github.com/LiveSchedule-v2/models"
	"github.com/gin-gonic/gin"
)

func RegisterDiscordId(c *gin.Context) {
	// Get the request
	userId, err := c.Cookie("UserId")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	var body struct {
		DiscordId int `json:"discord_id" form:"discord_id"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	user := models.User{}
	initializers.DB.First(&user, "ID = ?", userId)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	initializers.DB.Model(&user).Update("discord_id", body.DiscordId)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully update user data",
		"data": gin.H{
			"discord_id": body.DiscordId,
		},
	})
}
