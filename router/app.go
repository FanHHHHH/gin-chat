package router

import (
	"gin-chat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", service.GetIndex)
	r.GET("/user/getUserList", service.GetUserList)
	return r
}
