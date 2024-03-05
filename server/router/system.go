package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/middleware"
	"server/service"
)

func LoadSystemRouter(e *gin.Engine) {

	r := e.Group("/system")
	r.Use(middleware.Cors())
	{
		r.GET("/menu", service.GetMenuList)
		r.GET("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "获取登录信息",
			})
		})
	}
}
