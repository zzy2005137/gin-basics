package api

import (
	"gin-example/pkg/util"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) {

	//获取参数
	username := c.Query("username")
	password := c.Query("password")

	//验证参数，检查用户是否存在...
	if username != "" && password != "" {
		//生成token
		token := util.GenerateJwt(username, password)
		c.JSON(200, gin.H{"token": token})
		return
	}

	c.JSON(400, gin.H{"message": "缺少参数"})

}
