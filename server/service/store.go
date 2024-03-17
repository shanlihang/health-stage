package service

import (
	"server/dao"
	"strconv"

	"github.com/gin-gonic/gin"
)

type insertGoods struct {
	name   string
	remark string
	divide string
}

func GetGoodsList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"list": dao.SelectGoods(),
	})
}

func DropGoods(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("ID"), 10, 64)
	if err != nil {
		return
	}
	ctx.JSON(200, gin.H{
		"result": dao.DeleteGoods(id),
	})
}

func AddGoods(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"result": dao.InsertGoods(),
	})
}
