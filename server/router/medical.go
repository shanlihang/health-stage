package router

import (
	"server/service"

	"github.com/gin-gonic/gin"
)

func LoadMedicalRouter(e *gin.Engine) {
	r := e.Group("/medical")
	{
		r.GET("/community", service.GetCommunityList)

		r.POST("/community", service.AddCommunity)
		r.DELETE("/community", service.DropCommunity)
		r.GET("/result", service.GetResultList)
		r.POST("/result", service.AddResult)
		r.DELETE("/result", service.DropResult)
	}
}
