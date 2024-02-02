package main

import (
	"github.com/LiveSchedule-v2/controllers"
	"github.com/LiveSchedule-v2/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDb()
}

func main() {
	router := gin.Default()
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	router.Run()
}
