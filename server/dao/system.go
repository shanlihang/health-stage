package dao

import (
	"server/global"
	"server/model"
)

func SelectTopMenu() *[]model.TopMenu {
	menu := &[]model.TopMenu{}
	global.DB.Find(menu)
	return menu
}
