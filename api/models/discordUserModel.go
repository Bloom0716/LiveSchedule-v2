package models

import "gorm.io/gorm"

type DiscordUser struct {
	gorm.Model
	Email     string `gorm:"unique"`
	UserId    uint   `gorm:"unique"`
	DiscordId int    `gorm:"unique"`
	Name      string
	Password  string
}
