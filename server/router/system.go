package router

import (
	"net/http"
	"server/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func LoadSystemRouter(e *gin.Engine) {

	r := e.Group("/system")
	r.Use(cors.Default())
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
