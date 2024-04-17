package service

import (
	"github.com/gin-gonic/gin"
	"server/dao"
	"server/model"
	"strconv"
)

func AddResult(ctx *gin.Context) {
	res := &model.Result{}
	err := ctx.ShouldBindJSON(res)
	if err != nil {
		panic("绑定失败")
	}
	row := dao.InsertResult(res)
	if row != 0 {
		ctx.JSON(200, "插入成功")
	} else {
		ctx.JSON(200, "插入失败")
	}
}

func GetResultList(ctx *gin.Context) {
	ctx.JSON(200, dao.SelectResults())
}

func DropResult(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("ID"), 10, 64)
	if err != nil {
		panic("数据类型转换异常")
	}
	ctx.JSON(200, dao.DeleteResult(id))
}
