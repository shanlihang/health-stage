package dao

import (
	"server/global"
	"server/model"

	"gorm.io/gorm"
)

func SelectGoods() *[]model.Goods {
	goods := &[]model.Goods{}
	global.DB.Find(goods)
	return goods
}

func DeleteGoods(ID int64) *gorm.DB {
	result := global.DB.Where("ID = ?", ID).Delete(&model.Goods{})
	return result
}

func InsertGoods(name, remark, divide string, nums int) {
	goods := &model.Goods{
		Name:   name,
		Nums:   nums,
		Divide: divide,
		Remark: remark,
	}
	result := global.DB.Create(goods)
	if result.Error == nil {
		panic("error")
	}
}
