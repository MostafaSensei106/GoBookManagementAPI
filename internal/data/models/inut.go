package models

import (
	"log"

	"github.com/MostafaSensei106/GoBookManagementAPI/internal/config"
)

func Init() {
	config.Connect()
	db := config.GetDB()
	if err := db.AutoMigrate(&Author{}, &BoolCategory{}, &Book{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
}
