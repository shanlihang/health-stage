package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Nickname string `gorm:"column:nickname"`
	Sex      int    `gorm:"column:sex"`
	Phone    string `gorm:"column:phone"`
	mail     string `gorm:"column:mail"`
}

func (u Users) TableName() string {
	return "users"
}
