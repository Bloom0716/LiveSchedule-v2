package controllers

import (
	"net/http"

	"github.com/LiveSchedule-v2/initializers"
	"github.com/LiveSchedule-v2/models"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	// Get the req
	userId, err := c.Cookie("UserId")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Look up requested user
	user := models.User{}
	initializers.DB.First(&user, "ID = ?", userId)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	// Delete user
	initializers.DB.Delete(&user)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully delete user",
	})
}
