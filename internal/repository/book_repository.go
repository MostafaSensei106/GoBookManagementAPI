package repository

import (
	"github.com/jinzhu/gorm"

	models "github.com/MostafaSensei106/GoBookManagementAPI/internal/data/models/book_model"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	db.AutoMigrate(&models.Book{}, &models.Author{}, &models.Category{})
	return &BookRepository{db: db}
}

func (r *BookRepository) CreateBook(b *models.Book) (*models.Book, error) {
	err := r.db.Create(b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *BookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) GetBookByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.Preload("Author").Preload("Categories").First(&book, id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) UpdateBook(b *models.Book) (*models.Book, error) {
	err := r.db.Save(b).Error
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *BookRepository) DeleteBook(id uint) error {

	err := r.db.Delete(&models.Book{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
