package dao

import (
	"gorm.io/gorm"
	"server/global"
	"server/model"
)

func InsertResult(res *model.Result) int64 {
	result := global.DB.Create(res)
	return result.RowsAffected
}

func SelectResults() *[]model.Result {
	results := &[]model.Result{}
	global.DB.Find(results)
	return results
}

func DeleteResult(ID int64) *gorm.DB {
	result := global.DB.Delete(&model.Result{}, ID)
	return result
}
