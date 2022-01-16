package routers

import (
	"gin-example/middleware/auth"
	"gin-example/routers/api"
	v1 "gin-example/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("token", api.GetToken)

	apiv1 := router.Group("/api/v1")
	apiv1.Use(auth.CheckAuth())
	{
		apiv1.GET("/test", v1.TestHanler)

	}

	return router
}
