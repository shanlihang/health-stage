package middleware

import "github.com/gin-gonic/gin"

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("front")
		c.Next()
		println("end")
	}
}
