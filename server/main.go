package main

import (
	"net/http"
	"server/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r1 := r.Group("/")
	{
		r1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "123456",
			})
		})
	}
	router.LoadSystemRouter(r)
	router.LoadMedicalRouter(r)
	r.Run(":9090")
}
