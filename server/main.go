package main

import (
	"net/http"
	"server/global"
	"server/router"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/wldz?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	global.DB = db

	r := gin.Default()
	r.Use(cors.Default())
	r1 := r.Group("/")
	{
		r1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": global.SUCCESS,
				"msg":  global.GetMsg(global.SUCCESS),
			})
		})
	}
	router.LoadSystemRouter(r)
	router.LoadMedicalRouter(r)
	router.LoadStoreRouter(r)
	r.Run(":9090")
}
