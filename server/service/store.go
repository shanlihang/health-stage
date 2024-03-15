package service

import (
	"server/dao"

	"github.com/gin-gonic/gin"
)

func GetGoodsList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"list": dao.SelectGoods(),
	})
}

func DropGoods(ctx *gin.Context) {
	id := ctx.Param("ID")
	println(id)
	ctx.JSON(200, gin.H{
		"result": dao.DeleteGoods(id),
	})
}

func AddGoods(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "插入成功",
	})
}
