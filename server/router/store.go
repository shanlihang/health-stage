package router

import (
	"net/http"
	"server/service"

	"github.com/gin-gonic/gin"
)

func LoadStoreRouter(e *gin.Engine) {

	r := e.Group("/store")
	r.GET("/goods", service.GetGoodsList)
	r.POST("/addGoods", service.AddGoods)
	r.DELETE("/deleteGoods", service.DropGoods)
	r.GET("/a", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "bbb",
		})
	})
}
