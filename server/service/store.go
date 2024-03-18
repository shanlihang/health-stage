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

	// name := ctx.Query("name")
	// remark := ctx.Query("remark")
	// divide := ctx.Query("divide")
	// nums, err := strconv.Atoi((ctx.Query("nums")))
	// if err != nil {
	// 	panic("错误：")
	// }

	// name, remark, divide string, nums int
	ctx.JSON(200, gin.H{
		// "result": dao.InsertGoods(name, remark, divide, nums),
		"result": "1230",
	})
}
