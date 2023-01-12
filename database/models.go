package database

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id       uint64 `gorm:"primaryKey;autoIncrement:true"`
	Login    string
	Password string
}

type VKSession struct {
	gorm.Model
	Id           uint64 `gorm:"primaryKey;autoIncrement:true"`
	User_id      uint64
	VkId         uint64
	Access_token string
	Expires      time.Time
}

type AuthSession struct {
	gorm.Model
	Id            uint64 `gorm:"primaryKey;autoIncrement:true"`
	Access_token  string
	Refresh_token string
	User_id       uint64
	Session_start time.Time
}
