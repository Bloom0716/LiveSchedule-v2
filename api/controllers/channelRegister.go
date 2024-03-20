package controllers

import (
	"net/http"
	"strconv"

	"github.com/LiveSchedule-v2/initializers"
	"github.com/LiveSchedule-v2/models"
	"github.com/gin-gonic/gin"
)

func RegisterChannel(c *gin.Context) {
	// Get the request
	userIdStr, err := c.Cookie("UserId")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	var body struct {
		ChannelTitle string `form:"channel_title"`
		ChannelId    string `form:"channel_id"`
		ThumbnailUrl string `form:"thumbnail_url"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	if body.ChannelTitle == "" || body.ChannelId == "" || body.ThumbnailUrl == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unexpected error occured",
		})
		return
	}

	// Look up request channel
	checkChannel := models.Channel{}
	initializers.DB.Where("user_id = ? AND channel_id = ?", userId, body.ChannelId).First(&checkChannel)

	if checkChannel.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Already registerd this channel",
		})
		return
	}

	// Register channel
	channel := models.Channel{
		UserId:       uint(userId),
		ChannelTitle: body.ChannelTitle,
		ChannelId:    body.ChannelId,
		ThumbnailUrl: body.ThumbnailUrl,
	}
	result := initializers.DB.Create(&channel)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to register channel",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully register channel",
	})
}
