package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type errorResp struct {
	success   bool
	message   string
	timeStamp time.Time
}

type successResp struct {
	message string
	success bool
	data    interface{}
}

// Reusable err response function
func ErrorResponse(w http.ResponseWriter, message string, status int) {
	response := errorResp{
		message:   message,
		timeStamp: time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

// Reusable success response function
func SuccessResponse(w http.ResponseWriter, status int, data interface{}) {
	result := successResp{
		success: true,
		data:    data,
	}

	response, err := json.Marshal(result)

	if err != nil {
		return
	}

	log.Println(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
