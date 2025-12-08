package models

import (
	"log"

	"gorm.io/gorm"

	"github.com/MostafaSensei106/GoBookManagementAPI/internal/config"
)

var db *gorm.DB

func Init() {
	config.Connect()
	db := config.GetDB()
	if err := db.AutoMigrate(&Author{}, &BookCategory{}, &Book{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
}
