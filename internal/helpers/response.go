package helpers

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
	Data       T      `json:"data,omitempty"`
	Error      any    `json:"errors,omitempty"`
}

// New Response Maker
func NewResponse[T any](w http.ResponseWriter, statusCode int, msg string, data T, errors any) {
	res := Response[T]{
		StatusCode: statusCode,
		Message:    msg,
		Data:       data,
		Error:      errors,
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
