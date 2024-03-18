package dao

import (
	"server/global"
	"server/model"
)

func SelectAllUsers() *[]model.Users {
	users := &[]model.Users{}
	global.DB.Find(users)
	return users
}
