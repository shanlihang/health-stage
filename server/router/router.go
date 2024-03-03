package router

import (
	"github.com/gin-gonic/gin"
	"server/middleware"
)

func LoadRouters() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r1 := r.Group("/system")
	{
		r1.GET("/getMenu", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"name": "Lily",
				"age":  25,
				"sex":  "男",
			})
		})
	}
	r.Run(":9090")
}
