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

func (a *Author) CreateAuthor() (*Author, error) {
	if err := db.Create(a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func GetAllAuthors() ([]Author, error) {
	var authors []Author
	if err := db.Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func GetAuthorByID(id int64) (*Author, error) {
	var author Author
	result := db.First(&author, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &author, nil
}

func DeleteAuthor(id int64) error {
	var author Author
	result := db.First(&author, id)
	if result.Error != nil {
		return result.Error
	}
	return db.Delete(&author).Error
}

func UpdateAuthor(id int64, author *Author) error {
	result := db.First(&author, id)
	if result.Error != nil {
		return result.Error
	}
	return db.Save(&author).Error
}
