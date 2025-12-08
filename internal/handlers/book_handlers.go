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

// ---------------- Helper Functions ----------------

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

func parseID(r *http.Request, param string) (int64, error) {
	vars := mux.Vars(r)
	idStr := vars[param]
	return strconv.ParseInt(idStr, 10, 64)
}

// ---------------- Handlers ----------------

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to get books"})
		return
	}
	writeJSON(w, http.StatusOK, books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r, "book_id")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid book id"})
		return
	}

	book, err := models.GetBookByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "book not found"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to get book"})
		return
	}

	writeJSON(w, http.StatusOK, book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	if err := utils.ParseBody(r, book); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	b, err := book.CreateBook()
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to create book"})
		return
	}

	writeJSON(w, http.StatusOK, b)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r, "book_id")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid book id"})
		return
	}

	var fields map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&fields); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	updatedBook, err := models.UpdateBook(id, fields)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "book not found"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to update book"})
		return
	}

	writeJSON(w, http.StatusOK, updatedBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r, "book_id")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid book id"})
		return
	}

	err = models.DeleteBook(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "book not found"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "failed to delete book"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "book deleted successfully"})
}
