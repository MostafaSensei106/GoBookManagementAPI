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
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`

	Page  int64       `json:"page,omitempty"`
	Limit int         `json:"limit,omitempty"`
	Total int64       `json:"total,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

// ---------------- Helper Functions ----------------

// writeResponse: Unified JSON response (supports pagination automatically)
func writeResponse(w http.ResponseWriter, status int, message string, data interface{}, detail string,
	page int64, limit int, total int64) {

	resp := APIResponse{
		Code:    status,
		Message: message,
		Detail:  detail,
		Data:    data,
	}

	// Only show pagination fields if values were passed
	if page > 0 && limit > 0 {
		resp.Page = page
		resp.Limit = limit
		resp.Total = total
	}

	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

// Parse URL ID
func parseID(r *http.Request, param string) (int64, error) {
	vars := mux.Vars(r)
	idStr := vars[param]
	return strconv.ParseInt(idStr, 10, 64)
}

// ---------------- Handlers ----------------

// Get All Books + Pagination
func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 20
	}

	books, total, err := models.GetBooksPaginated(page, limit)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, "Failed to get books", nil, err.Error(), 0, 0, 0)
		return
	}

	writeResponse(w, http.StatusOK, "Books retrieved successfully", books, "", int64(page), limit, total)
}

// Get Book By ID
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r, "id")
	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid book ID", nil, err.Error(), 0, 0, 0)
		return
	}

	book, err := models.GetBookByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeResponse(w, http.StatusNotFound, "Book not found", nil, err.Error(), 0, 0, 0)
			return
		}
		writeResponse(w, http.StatusInternalServerError, "Failed to get book", nil, err.Error(), 0, 0, 0)
		return
	}

	writeResponse(w, http.StatusOK, "Book retrieved successfully", book, "", 0, 0, 0)
}

// Create Book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	if err := utils.ParseBody(r, book); err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid request body", nil, err.Error(), 0, 0, 0)
		return
	}

	if err := book.Validate(); err != nil {
		writeResponse(w, http.StatusBadRequest, "Validation failed", nil, err.Error(), 0, 0, 0)
		return
	}

	b, err := book.CreateBook()
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, "Failed to create book", nil, err.Error(), 0, 0, 0)
		return
	}

	writeResponse(w, http.StatusOK, "Book created successfully", b, "", 0, 0, 0)
}

// Update Book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r, "id")
	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid book ID", nil, err.Error(), 0, 0, 0)
		return
	}

	var fields map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&fields); err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid request body", nil, err.Error(), 0, 0, 0)
		return
	}

	updatedBook, err := models.UpdateBook(id, fields)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeResponse(w, http.StatusNotFound, "Book not found", nil, err.Error(), 0, 0, 0)
			return
		}
		writeResponse(w, http.StatusInternalServerError, "Failed to update book", nil, err.Error(), 0, 0, 0)
		return
	}

	writeResponse(w, http.StatusOK, "Book updated successfully", updatedBook, "", 0, 0, 0)
}

// Delete Book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r, "id")
	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid book ID", nil, err.Error(), 0, 0, 0)
		return
	}

	err = models.DeleteBook(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeResponse(w, http.StatusNotFound, "Book not found", nil, err.Error(), 0, 0, 0)
			return
		}
		writeResponse(w, http.StatusInternalServerError, "Failed to delete book", nil, err.Error(), 0, 0, 0)
		return
	}

	writeResponse(w, http.StatusOK, "Book deleted successfully", nil, "", 0, 0, 0)
}
