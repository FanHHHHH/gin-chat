package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} json "{"code":200,"data":"welcome"}"
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "welcome",
	})
}

// func GetUserList(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"messaga":
// 	})
// }
