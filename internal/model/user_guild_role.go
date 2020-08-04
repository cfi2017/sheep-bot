package model

import "github.com/jinzhu/gorm"

type UserGuildRole struct {
	gorm.Model
	UserID  string
	RoleID  string
	GuildID string
}
