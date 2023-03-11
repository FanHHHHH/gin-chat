package service

import (
	"gin-chat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags User
// @Success 200 {string} json{"code","data"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	users := make([]*models.UserBasics, 10)
	users = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
