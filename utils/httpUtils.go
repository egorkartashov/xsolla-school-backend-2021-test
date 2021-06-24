package utils

import (
	"encoding/json"
	"net/http"
)

func RespondJson(writer http.ResponseWriter, statusCode int, data map[string]interface{}) {
	writer.WriteHeader(statusCode)
	writer.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(data)
}
