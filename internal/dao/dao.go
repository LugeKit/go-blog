package dao

import "gorm.io/gorm"

var db *gorm.DB

const (
	dbConnectString = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local"
)
