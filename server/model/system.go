package model

import "gorm.io/gorm"

type TopMenu struct {
	gorm.Model
	Name   string `gorm:"column:name" json:"name"`
	Icon   string `gorm:"column:icon" json:"icon"`
	Router string `gorm:"column:router" json:"router"`
}

func (t TopMenu) TableName() string {
	return "top_menu"
}
