package controllers

import (
	"net/http"

	"github.com/LiveSchedule-v2/initializers"
	"github.com/LiveSchedule-v2/models"
	"github.com/gin-gonic/gin"
)

func DeleteChannel(c *gin.Context) {
	userId, err := c.Cookie("UserId")
	channelId := c.Query("channelId")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Look up requested channel
	channel := models.Channel{}
	initializers.DB.Where("user_id = ? AND channel_id = ?", userId, channelId).First(&channel)

	if channel.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	// Delete channel
	initializers.DB.Delete(&channel)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully delete channel",
	})
}
