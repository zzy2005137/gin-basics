package auth

import (
	"gin-example/pkg/util"

	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get parameters
		token := c.Query("token")
		//validate token
		//可能的情况有无效或过期，返回具体的错误码和信息（未完成）
		_, err := util.ParseToken(token)
		if err != nil {
			c.JSON(400, gin.H{"error": err})
			c.Abort()
			return
		} else {
			c.Next()
		}

	}
}
