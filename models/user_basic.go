package models

import (
	"fmt"
	"gin-chat/utils"
	"time"

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
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	IsOnline      bool
	DeviceInfo    string
}

func (table *UserBasics) TableName() string {
	return "user_basics"
}

func GetUserList() []*UserBasics {
	users := make([]*UserBasics, 10)
	utils.DB.Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
	return users
}
