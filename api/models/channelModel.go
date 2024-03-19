package models

import "gorm.io/gorm"

type Channel struct {
	gorm.Model
	UserId       uint
	ChannelTitle string
	ChannelId    string
	ThumbnailUrl string
}
