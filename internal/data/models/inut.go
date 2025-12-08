package models

import (
	"log"

	"github.com/MostafaSensei106/GoBookManagementAPI/internal/config"
)

func init() {
	config.Connect()
	db = config.GetDB()
	if err := db.AutoMigrate(&Book{}, &Author{}, &Category{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
}
