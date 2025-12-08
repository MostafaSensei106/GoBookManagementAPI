package models

import "gorm.io/gorm"

// -------------------- BoolCategory --------------------
type BoolCategory struct {
	gorm.Model
	Name  string `json:"name"`
	Books []Book `gorm:"many2many:book_categories;" json:"books"`
}

func (b *BoolCategory) CreateCategory() (*BoolCategory, error) {
	if err := db.Create(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func GetAllCategories() ([]BoolCategory, error) {
	var categories []BoolCategory
	if err := db.Preload("Books").
		Find(&categories).
		Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoryByID(id int64) (*BoolCategory, error) {
	var category BoolCategory
	result := db.Preload("Books").First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func DeleteCategory(id int64) error {
	var category BoolCategory
	result := db.First(&category, id)
	if result.Error != nil {
		return result.Error
	}
	return db.Delete(&category).Error
}

func UpdateCategory(id int64, category *BoolCategory) error {
	result := db.First(&category, id)
	if result.Error != nil {
		return result.Error
	}
	return db.Save(&category).Error
}
