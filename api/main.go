package main

import (
	"github.com/LiveSchedule-v2/controllers"
	"github.com/LiveSchedule-v2/initializers"
	"github.com/LiveSchedule-v2/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDb()
}

func main() {
	router := gin.Default()

	userRouter := router.Group("users")
	{
		userRouter.POST("signup", controllers.Signup)
		userRouter.POST("login", controllers.Login)
		userRouter.GET("user", middleware.RequireAuth, controllers.GetUser)
		userRouter.PATCH("user", middleware.RequireAuth, controllers.UpdateUser)
		userRouter.DELETE("delete", middleware.RequireAuth, controllers.DeleteUser)
	}

	channelRouter := router.Group("channels")
	{
		channelRouter.GET("search", controllers.SearchChannel)
		channelRouter.GET("channel", controllers.GetChannel)
		channelRouter.POST("register", controllers.RegisterChannel)
		channelRouter.DELETE("delete", controllers.DeleteChannel)
	}

	videoRouter := router.Group("videos")
	{
		videoRouter.GET("search", controllers.SearchVideo)
	}

	discordRouter := router.Group("discord")
	{
		discordRouter.POST("login", controllers.LoginDiscord)
		discordRouter.GET("search", controllers.GetDiscordVideo)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	router.Run()
}
