package config

import (
	"github.com/jinzhu/gorm"

	"github.com/MostafaSensei106/GoBookManagementAPI/internal/constants"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(constants.DataBaseClient, "root:root@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d

}

func GetDB() *gorm.DB {
	return db
}
