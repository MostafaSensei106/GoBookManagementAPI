package models

import "github.com/jinzhu/gorm"

// -------------------- Category --------------------
type Category struct {
	gorm.Model
	Name string `json:"name"`

	Books []Book `gorm:"many2many:book_categories;" json:"books"`
}
