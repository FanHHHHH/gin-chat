package service

import (
	"gin-chat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	users := make([]*models.UserBasics, 10)
	users = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"message": users,
	})
}
