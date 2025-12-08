package models

import (
	"github.com/jinzhu/gorm"
)

// -------------------- Category --------------------
type Category struct {
	gorm.Model
	Name string `json:"name"` // GORM هيعمل column name لوحده

	// علاقة Many-to-Many مع Book
	Books []Book `gorm:"many2many:book_categories;" json:"books"`
}

// -------------------- Author --------------------
type Author struct {
	gorm.Model
	Name  string `json:"name"`
	Bio   string `json:"bio"`
	Books []Book `gorm:"foreignKey:AuthorID" json:"books"`
}

// -------------------- Book --------------------
type Book struct {
	gorm.Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ISBN        string     `json:"isbn"`
	AuthorID    uint       `json:"author_id"`
	Author      Author     `gorm:"foreignKey:AuthorID" json:"author"`
	Categories  []Category `gorm:"many2many:book_categories;" json:"categories"`
	Publication string     `json:"publication"`
}
