package helpers

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse[T any] struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status    int    `json:"status"`
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	ErrorType string `json:"error_type,omitempty"`
	Errors    any    `json:"errors,omitempty"`
}

// New Response Maker
func NewSuccessResponse[T any](w http.ResponseWriter, statusCode int, msg string, data T, errors any) {
	res := SuccessResponse[T]{
		Status:  statusCode,
		Success: true,
		Message: msg,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

func NewErrorResponse(w http.ResponseWriter, statusCode int, msg string, error_type string, errors any) {
	res := ErrorResponse{
		Status:    statusCode,
		Success:   false,
		Message:   msg,
		ErrorType: error_type,
		Errors:    errors,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

// func NewResponse[T any](statusCode int, msg string, data T) Response[T] {
// 	return Response[T]{
// 		Code:    statusCode,
// 		Message: msg,
// 		Data:    data,
// 	}

// }

// Write JSON response
// func SendJSON(w http.ResponseWriter, statusCode int, res *Response[any]) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(statusCode)
// 	json.NewEncoder(w).Encode(res)
// }
