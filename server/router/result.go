package router

import (
	"github.com/gin-gonic/gin"
	"server/service"
)

func LoadResultRouter(e *gin.Engine) {
	r := e.Group("/result")
	{
		r.GET("/result", service.GetCommunityList)
		r.POST("/result", service.AddCommunity)
		r.DELETE("/result", service.DropCommunity)
	}
}
