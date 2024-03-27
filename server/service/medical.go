package service

import (
	"github.com/gin-gonic/gin"
	"server/dao"
	"server/model"
	"strconv"
)

func AddCommunity(ctx *gin.Context) {
	comm := &model.Community{}
	err := ctx.ShouldBindJSON(comm)
	if err != nil {
		panic("绑定失败")
	}
	row := dao.InsertCommunity(comm)
	println(row)
}

func GetCommunityList(ctx *gin.Context) {
	ctx.JSON(200, dao.SelectCommunities())
}

func DropCommunity(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("ID"), 10, 64)
	if err != nil {
		panic("数据类型转换异常")
	}
	ctx.JSON(200, dao.DeleteCommunity(id))
}
