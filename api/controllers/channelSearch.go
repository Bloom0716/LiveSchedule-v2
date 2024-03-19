package controllers

import (
	"net/http"
	"os"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"

	"github.com/gin-gonic/gin"
)

func SearchChannel(c *gin.Context) {
	// Get API Key
	apiKey := os.Getenv("API_KEY")

	// Youtube Data API
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

	// Search channel
	searchQuery := c.Query("q")
	if searchQuery == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing 'q' parameter",
		})
		return
	}

	searchCall := service.Search.List([]string{"snippet"}).Q(searchQuery).Type("channel").MaxResults(1)

	response, err := searchCall.Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to search channel",
		})
		return
	}

	channels := []gin.H{}
	for _, item := range response.Items {
		channel := gin.H{
			"title":     item.Snippet.Title,
			"channelId": item.Snippet.ChannelId,
			"thumbnail": item.Snippet.Thumbnails.Default.Url,
		}
		channels = append(channels, channel)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully search channels",
		"data":    channels,
	})
}
