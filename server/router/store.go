package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/middleware"
	"server/service"
)

func LoadStoreRouter(e *gin.Engine) {

	r := e.Group("/store")
	r.Use(middleware.Cors())
	{
		r.GET("/goods", service.GetGoodsList)
		r.POST("/addGoods", service.AddGoods)
		r.GET("/a", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "bbb",
			})
		})
	}
}
