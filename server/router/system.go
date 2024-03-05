package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadSystemRouter(e *gin.Engine) {

	r := e.Group("/system")
	{
		r.GET("/menu", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "获取menu信息",
			})
		})
		r.GET("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "获取登录信息",
			})
		})
	}
}
