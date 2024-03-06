package service

import (
	"github.com/gin-gonic/gin"
	"server/dao"
	"strconv"
)

func GetGoodsList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"list": dao.SelectGoods(),
	})
}

func AddGoods(ctx *gin.Context) {
	name := ctx.PostForm("name")
	nums, err := strconv.Atoi(ctx.PostForm("nums"))
	devide := ctx.PostForm("uint")
	remark := ctx.PostForm("remark")
	if err != nil {
		panic("转换整数失败")
		return
	}
	dao.InsertGoods(name, remark, devide, nums)
	ctx.JSON(200, gin.H{
		"msg": "插入成功",
	})
}
