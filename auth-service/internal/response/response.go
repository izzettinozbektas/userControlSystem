package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, statusCode int, res Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

func Success(w http.ResponseWriter, message string, data interface{}) {
	JSON(w, http.StatusOK, Response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func Fail(w http.ResponseWriter, statusCode int, message string, err interface{}) {
	JSON(w, statusCode, Response{
		Status:  false,
		Message: message,
		Error:   err,
	})
}
