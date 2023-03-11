package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome",
	})
}

// func GetUserList(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"messaga":
// 	})
// }
