package service

import (
	"server/dao"

	"github.com/gin-gonic/gin"
)

func GetMenusList(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"list": dao.SelectTopMenu(),
	})
}
