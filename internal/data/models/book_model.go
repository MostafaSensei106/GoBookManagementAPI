package models

import (
	"errors"

	"gorm.io/gorm"
)

// -------------------- Book --------------------
type Book struct {
	gorm.Model
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ISBN        string         `json:"isbn"`
	AuthorID    uint           `json:"author_id"`
	Author      Author         `gorm:"foreignKey:AuthorID" json:"author"`
	Categories  []BookCategory `gorm:"many2many:books_categories;" json:"categories"`
	Publication string         `json:"publication"`
}

// ---------------- Validation ----------------
func (b *Book) Validate() error {
	if b.Name == "" {
		return errors.New("name cannot be empty")
	}
	if b.Description == "" {
		return errors.New("description cannot be empty")
	}
	if b.ISBN == "" {
		return errors.New("ISBN cannot be empty")
	}
	if b.AuthorID == 0 {
		return errors.New("author_id cannot be empty")
	}
	if b.Publication == "" {
		return errors.New("publication cannot be empty")
	}
	return nil
}

// ---------------- Copy With ----------------
func (b *Book) copyWith(fields map[string]interface{}) *Book {
	newBook := *b
	for key, value := range fields {
		switch key {
		case "name":
			if v, ok := value.(string); ok {
				newBook.Name = v
			}
		case "description":
			if v, ok := value.(string); ok {
				newBook.Description = v
			}
		case "isbn":
			if v, ok := value.(string); ok {
				newBook.ISBN = v
			}
		case "author_id":
			if v, ok := value.(uint); ok {
				newBook.AuthorID = v
			}
		case "categories":
			if v, ok := value.([]BookCategory); ok {
				newBook.Categories = v
			}
		case "publication":
			if v, ok := value.(string); ok {
				newBook.Publication = v
			}
		}
	}
	return &newBook
}

// ---------------- CRUD ----------------
func (b *Book) CreateBook() (*Book, error) {
	if err := b.Validate(); err != nil {
		return nil, err
	}
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

func UpdateBook(id int64, fields map[string]interface{}) (*Book, error) {
	var book Book
	if err := db.Preload("Author").Preload("Categories").First(&book, id).Error; err != nil {
		return nil, err
	}
	updatedBook := book.copyWith(fields)
	if err := updatedBook.Validate(); err != nil {
		return nil, err
	}
	if err := db.Save(updatedBook).Error; err != nil {
		return nil, err
	}
	return updatedBook, nil
}
