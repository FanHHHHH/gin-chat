package models

import (
	"gorm.io/gorm"
)

type UserBasics struct {
	gorm.Model
	Identity      string
	Name          string
	Password      string
	Phone         string
	Email         string
	Description   string
	ClientIp      string
	ClientPort    string
	LoginTime     uint64
	HeartbeatTime uint64
	LoginOutTime  uint64
	IsOnline      bool
	DeviceInfo    string
}

func (table *UserBasics) TableName() string {
	return "user_basics"
}
