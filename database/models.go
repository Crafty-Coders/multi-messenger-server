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
	Id          uint64 `gorm:"primaryKey;autoIncrement:true"`
	UserId      uint64
	VkId        uint64
	AccessToken string
	Expires     time.Time
}

type AuthSession struct {
	gorm.Model
	Id           uint64 `gorm:"primaryKey;autoIncrement:true"`
	AccessToken  string
	RefreshToken string
	UserId       uint64
	SessionStart time.Time
}
