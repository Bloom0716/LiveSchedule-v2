package controllers

import (
	"net/http"

	"github.com/LiveSchedule-v2/initializers"
	"github.com/LiveSchedule-v2/models"
	"github.com/gin-gonic/gin"
)

func GetChannel(c *gin.Context) {
	// Get the params
	userId, err := c.Cookie("UserId")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Look up requested userId
	user := models.User{}
	initializers.DB.First(&user, "ID = ?", userId)

	if user.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	var channels []models.Channel
	initializers.DB.Where("user_id = ?", userId).Find(&channels)

	data := []gin.H{}
	for _, item := range channels {
		data_item := gin.H{
			"title":      item.ChannelTitle,
			"channel_id": item.ChannelId,
			"thumbnail":  item.ThumbnailUrl,
		}
		data = append(data, data_item)
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully get channels",
		"data":    data,
	})
}
