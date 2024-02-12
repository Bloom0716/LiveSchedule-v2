package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/LiveSchedule-v2/initializers"
	"github.com/LiveSchedule-v2/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get token
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Require ot invalid token",
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "This token is expired",
			})
			return
		}

		// Find the user with token sub
		user := models.User{}
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}

		c.Set("user", user)

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}

	// Continue
	c.Next()
}
