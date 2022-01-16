package v1

import "github.com/gin-gonic/gin"

func TestHanler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "test"})
}
