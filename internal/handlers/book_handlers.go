package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"github.com/MostafaSensei106/GoBookManagementAPI/internal/constants"
	"github.com/MostafaSensei106/GoBookManagementAPI/internal/data/models"
	"github.com/MostafaSensei106/GoBookManagementAPI/internal/utils"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	var err error

	books, err = models.GetAllBooks()
	if err != nil {
		panic(err)
	}

	res, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}

	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["book_id"]

	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	bookDetails, err := models.GetBookByID(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "failed to marshal book details", http.StatusInternalServerError)
		return
	}

	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	if err := json.NewEncoder(w).Encode(bookDetails); err != nil {
		http.Error(w, "failed to encode book", http.StatusInternalServerError)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)

	b, err := CreateBook.CreateBook()
	if err != nil {
		http.Error(w, "failed to create book", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(b)
	if err != nil {
		http.Error(w, "failed to marshal book", http.StatusInternalServerError)
		return
	}
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(constants.ContentType, constants.ApplicationJson)

	vars := mux.Vars(r)
	idstr := vars["book_id"]

	ID, err := strconv.ParseInt(idstr, 0, 0)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}
	var faileds map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&faileds); err != nil {
		http.Error(w, "failed to decode book", http.StatusInternalServerError)
		return
	}
	updatedBook, err := models.UpdateBook(ID, faileds)
	if err != nil {
		http.Error(w, "failed to update book", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(updatedBook)
	if err != nil {
		http.Error(w, "failed to marshal book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["book_id"]

	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	err = models.DeleteBook(ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "book not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to delete book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "book deleted successfully"}`))

}
