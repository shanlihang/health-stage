package dao

import (
	"gorm.io/gorm"
	"server/global"
	"server/model"
)

func InsertCommunity(comm *model.Community) int64 {
	result := global.DB.Create(comm)
	return result.RowsAffected
}

func SelectCommunities() *[]model.Community {
	communities := &[]model.Community{}
	global.DB.Find(communities)
	return communities
}

func DeleteCommunity(ID int64) *gorm.DB {
	result := global.DB.Delete(&model.Community{}, ID)
	return result
}
