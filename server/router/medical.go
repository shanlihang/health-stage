package router

import (
	"net/http"
	"server/service"

	"github.com/gin-gonic/gin"
)

func LoadMedicalRouter(e *gin.Engine) {
	r := e.Group("/medical")
	{
		r.GET("/community", service.GetCommunityList)

		r.POST("/community", service.AddCommunity)
		r.DELETE("/community", service.DropCommunity)
		r.GET("/result", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "获取结果信息",
			})
		})
	}
}
