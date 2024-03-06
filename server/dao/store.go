package dao

import (
	"server/global"
	"server/model"
)

func SelectGoods() *[]model.Goods {
	goods := &[]model.Goods{}
	global.DB.Find(goods)
	return goods
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
