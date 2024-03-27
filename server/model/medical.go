package model

import "gorm.io/gorm"

type Community struct {
	gorm.Model
	POIId       string `gorm:"column:POI_id" json:"id"`
	POIName     string `gorm:"column:POI_name" json:"name"`
	POIDistrict string `gorm:"column:POI_district" json:"district"`
	POIAdcode   string `gorm:"column:POI_adcode" json:"adcode"`
	POILocation string `gorm:"column:POI_location" json:"location"`
	POIAddress  string `gorm:"column:POI_address" json:"address"`
}

func (c Community) TableName() string {
	return "community"
}
