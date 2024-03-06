package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	name   string `gorm:"column:name"`
	nums   int    `gorm:"column:nums"`
	remark string `gorm:"column:remark"`
}

func (g Goods) TableName() string {
	return "goods"
}
