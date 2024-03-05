package router

import (
	"net/http"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func LoadMedicalRouter(e *gin.Engine) {
	r := e.Group("/medical")
	r.Use(middleware.Test())
	{
		r.GET("/community", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "获取社区信息",
			})
		})
		r.GET("/result", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "获取结果信息",
			})
		})
	}
}
