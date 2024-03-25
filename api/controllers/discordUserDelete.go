package controllers

import (
	"net/http"
	"strconv"

	"github.com/LiveSchedule-v2/initializers"
	"github.com/LiveSchedule-v2/models"
	"github.com/gin-gonic/gin"
)

func DeleteDiscordUser(c *gin.Context) {
	// Get the req
	discordIdStr := c.Query("discordId")
	discordId, _ := strconv.Atoi(discordIdStr)

	// Look up requested DiscordUser
	discordUser := models.DiscordUser{}
	initializers.DB.First(&discordUser, "discord_id = ?", discordId)

	if discordUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid discordId",
		})
		return
	}

	// Delete DiscordUser
	initializers.DB.Delete(&discordUser)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully delete discordUser",
	})
}
