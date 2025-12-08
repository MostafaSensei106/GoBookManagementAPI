package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/MostafaSensei106/GoBookManagementAPI/internal/constants"
	"github.com/MostafaSensei106/GoBookManagementAPI/internal/data/models"
	"github.com/MostafaSensei106/GoBookManagementAPI/internal/utils"
)

// ---------------- Response Struct ----------------

type APIResponse struct {
	Code    int         `json:"code"`             // HTTP status code
	Message string      `json:"message"`          // رسالة نجاح أو فشل
	Data    interface{} `json:"data,omitempty"`   // البيانات لو فيه
	Detail  string      `json:"detail,omitempty"` // تفاصيل إضافية للأخطاء أو الفالديشن
}

// ---------------- Helper Functions ----------------

func writeResponse(w http.ResponseWriter, status int, message string, data interface{}, detail string) {
	resp := APIResponse{
		Code:    status,
		Message: message,
		Data:    data,
		Detail:  detail,
	}
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
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
		writeResponse(w, http.StatusInternalServerError, "Failed to get books", nil, err.Error())
		return
	}
	writeResponse(w, http.StatusOK, "Books retrieved successfully", books, "")
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r, "id")
	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid book ID", nil, err.Error())
		return
	}

	book, err := models.GetBookByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeResponse(w, http.StatusNotFound, "Book not found", nil, err.Error())
			return
		}
		writeResponse(w, http.StatusInternalServerError, "Failed to get book", nil, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, "Book retrieved successfully", book, "")
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	if err := utils.ParseBody(r, book); err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid request body", nil, err.Error())
		return
	}

	if err := book.Validate(); err != nil {
		writeResponse(w, http.StatusBadRequest, "Validation failed", nil, err.Error())
		return
	}

	b, err := book.CreateBook()
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, "Failed to create book", nil, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, "Book created successfully", b, "")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r, "id")
	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid book ID", nil, err.Error())
		return
	}

	var fields map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&fields); err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid request body", nil, err.Error())
		return
	}

	updatedBook, err := models.UpdateBook(id, fields)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeResponse(w, http.StatusNotFound, "Book not found", nil, err.Error())
			return
		}
		if err := updatedBook.Validate(); err != nil { // validation check
			writeResponse(w, http.StatusBadRequest, "Validation failed", nil, err.Error())
			return
		}
		writeResponse(w, http.StatusInternalServerError, "Failed to update book", nil, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, "Book updated successfully", updatedBook, "")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r, "id")
	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid book ID", nil, err.Error())
		return
	}

	err = models.DeleteBook(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeResponse(w, http.StatusNotFound, "Book not found", nil, err.Error())
			return
		}
		writeResponse(w, http.StatusInternalServerError, "Failed to delete book", nil, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, "Book deleted successfully", nil, "")
}
