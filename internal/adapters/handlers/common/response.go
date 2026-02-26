// Package common provides shared utilities for the adapter layer,
// including standardized HTTP response formats and error mapping
// to ensure consistent API communication.
package common

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteSuccess(w http.ResponseWriter, status int, message string, data interface{}) {
	JSON(w, status, APIResponse{Success: true, Message: message, Data: data})
}

func WriteError(w http.ResponseWriter, status int, message string) {
	JSON(w, status, APIResponse{Success: false, Message: message})
}
