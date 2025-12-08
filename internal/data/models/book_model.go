package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// -------------------- Book --------------------
type Book struct {
	gorm.Model
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ISBN        string         `json:"isbn"`
	AuthorID    uint           `json:"author_id"`
	Author      Author         `gorm:"foreignKey:AuthorID" json:"author"`
	Categories  []BoolCategory `gorm:"many2many:book_categories;" json:"categories"`
	Publication string         `json:"publication"`
}

func (b *Book) CreateBook() (*Book, error) {
	if err := db.Create(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}
func GetAllBooks() ([]Book, error) {
	var books []Book
	if err := db.Preload("Author").Preload("Categories").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBookByID(id int64) (*Book, error) {
	var book Book
	result := db.Preload("Author").Preload("Categories").First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func DeleteBook(id int64) error {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		return result.Error
	}
	return db.Delete(&book).Error
}

func UpdateBook(id int64, book *Book) error {
	result := db.First(&book, id)
	if result.Error != nil {
		return result.Error
	}
	return db.Save(&book).Error
}
