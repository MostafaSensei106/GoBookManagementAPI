package models

import (
	"github.com/jinzhu/gorm"
)

// -------------------- Author --------------------
type Author struct {
	gorm.Model
	Name  string `json:"name"`
	Bio   string `json:"bio"`
	Books []Book `gorm:"foreignKey:AuthorID" json:"books"`
}
