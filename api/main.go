package main

import (
	"github.com/LiveSchedule-v2/controllers"
	"github.com/LiveSchedule-v2/initializers"
	"github.com/LiveSchedule-v2/middleware"
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
		userRouter.GET("/:userId", middleware.RequireAuth, controllers.GetUser)
		userRouter.PATCH("/:userId", middleware.RequireAuth, controllers.UpdateUser)
		userRouter.DELETE("/delete/:userId", middleware.RequireAuth, controllers.DeleteUser)
	}

	router.Run()
}
