package router

import (
	"server/service"

	"github.com/gin-gonic/gin"
)

func LoadUserRouter(e *gin.Engine) {

	r := e.Group("/user")
	r.GET("/list", service.GetUsersList)

}
