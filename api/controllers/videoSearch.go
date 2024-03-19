package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func SearchVideo(c *gin.Context) {
	// Get API key
	apiKey := os.Getenv("API_KEY")

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

	// Get upcoming videos
	channelId := c.Query("channelId")
	if channelId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Missing 'channelId' parameter",
		})
		return
	}

	// now := time.Now().Format(time.RFC3339)
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

	videos := []gin.H{}
	for _, item := range response.Items {
		video := gin.H{
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

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully search videos",
		"data":    videos,
	})
}
