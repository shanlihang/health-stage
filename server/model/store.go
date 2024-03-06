package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Name   string `gorm:"column:name"`
	Nums   int    `gorm:"column:nums"`
	Divide string `gorm:"column:uint"`
	Remark string `gorm:"column:remark"`
}

func (g Goods) TableName() string {
	return "goods"
}
