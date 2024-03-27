package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/LiveSchedule-v2/initializers"
	"github.com/LiveSchedule-v2/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func GetDiscordVideo(c *gin.Context) {
	// Get apiKey and discordId
	apiKey := os.Getenv("API_KEY")
	discordId := c.Query("discordId")

	// Look up requested DiscordUser
	user := models.User{}
	initializers.DB.First(&user, "discord_id = ?", discordId)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid User",
		})
		return
	}

	// YouTube Data API
	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create YouTube service: " + err.Error(),
		})
		return
	}

	// Get channelId
	channels := []models.Channel{}
	initializers.DB.Where("user_id = ?", user.ID).Find(&channels)

	// Get videos
	videos := []gin.H{}
	for _, channel := range channels {
		channelId := channel.ChannelId

		searchCall := service.Search.List([]string{"snippet"}).
			ChannelId(channelId).
			EventType("upcoming").
			Type("video").
			MaxResults(1).
			Order("date")

		response, err := searchCall.Do()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch videos",
			})
			return
		}

		for _, item := range response.Items {
			video := gin.H{
				"channelTitle":       channel.ChannelTitle,
				"channelId":          channelId,
				"title":              item.Snippet.Title,
				"videoId":            item.Id.VideoId,
				"thumbnail":          item.Snippet.Thumbnails.Default.Url,
				"scheduledStartTime": "",
			}

			if item.Snippet.LiveBroadcastContent == "upcoming" {
				videoDetails, err := service.Videos.List([]string{"liveStreamingDetails"}).Id(item.Id.VideoId).Do()
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "Failed to retrieve video details: " + err.Error(),
					})
					return
				}

				if len(videoDetails.Items) == 0 {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "There are no video data",
					})
					return
				}

				scheduledStartTime, err := time.Parse(time.RFC3339, videoDetails.Items[0].LiveStreamingDetails.ScheduledStartTime)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "Failed to parse scheduled start time: " + err.Error(),
					})
					return
				}
				video["scheduledStartTime"] = scheduledStartTime.Format(time.RFC3339)
			}
			videos = append(videos, video)
		}
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully search videos",
		"data":    videos,
	})
}
