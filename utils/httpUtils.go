package utils

import (
	"encoding/json"
	"net/http"
)

func RespondJson(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.WriteHeader(statusCode)
	writer.Header().Add("Content-Type", "application/json")
	if data != nil {
		_ = json.NewEncoder(writer).Encode(data)
	}
}

func RespondErrorJson(writer http.ResponseWriter, statusCode int, errorMessage string) {
	errorMessageBody := map[string]string{"error": errorMessage}
	RespondJson(writer, statusCode, errorMessageBody)
}
