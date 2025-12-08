package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MostafaSensei106/GoBookManagementAPI/internal/constants"
)

type NetworkError struct {
	StatusCode int    `json:"status_code"`
	ErrorCode  string `json:"error_code"`
	Message    string `json:"message"`
}

func jsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func jsonError(w http.ResponseWriter, err NetworkError) {
	w.Header().Set(constants.ContentType, constants.ApplicationJson)
	w.WriteHeader(err.StatusCode)
	json.NewEncoder(w).Encode(err)
}
