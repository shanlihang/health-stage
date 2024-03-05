package service

import (
	"github.com/gin-gonic/gin"
	"server/dao"
)

func GetMenuList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"list": dao.SelectTopMenu(),
	})
}
