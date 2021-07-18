package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"net/http"
	"strconv"
	"strings"
)

func RespondJson(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	if data != nil {
		_ = json.NewEncoder(writer).Encode(data)
	}
}

func RespondErrorJson(writer http.ResponseWriter, statusCode int, errorMessage string) {
	errorMessageBody := map[string]string{"error": errorMessage}
	RespondJson(writer, statusCode, errorMessageBody)
}

func RespondValidationErrors(writer http.ResponseWriter, errors validator.ValidationErrors) {
	builder := strings.Builder{}
	for _, e := range errors {
		builder.WriteString(fmt.Sprintf("%s\n", e))
	}
	RespondErrorJson(writer, http.StatusBadRequest, builder.String())
}

func TryParseIntQueryParameterOrDefault(request *http.Request, paramKey string, defaultValue int) (int, bool) {
	var paramValue = defaultValue
	var err error
	paramValueStr := request.URL.Query().Get(paramKey)

	if paramValueStr != "" {
		if paramValue, err = strconv.Atoi(paramValueStr); err == nil {
			return paramValue, true
		}

		return 0, false
	}

	return paramValue, true
}
