package models

import "gorm.io/gorm"

// -------------------- BookCategory --------------------
type BookCategory struct {
	gorm.Model
	Name  string `json:"name"`
	Books []Book `gorm:"many2many:books_categories;" json:"books"`
}

func (b *BookCategory) CreateCategory() (*BookCategory, error) {
	if err := db.Create(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func GetAllCategories() ([]BookCategory, error) {
	var categories []BookCategory
	if err := db.Preload("Books").
		Find(&categories).
		Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoryByID(id int64) (*BookCategory, error) {
	var category BookCategory
	result := db.Preload("Books").First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func DeleteCategory(id int64) error {
	var category BookCategory
	result := db.First(&category, id)
	if result.Error != nil {
		return result.Error
	}
	return db.Delete(&category).Error
}

func UpdateCategory(id int64, category *BookCategory) error {
	result := db.First(&category, id)
	if result.Error != nil {
		return result.Error
	}
	return db.Save(&category).Error
}
