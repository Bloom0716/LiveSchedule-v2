package initializers

import "github.com/LiveSchedule-v2/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Channel{})
	// DB.AutoMigrate(&models.DiscordUser{})
}
