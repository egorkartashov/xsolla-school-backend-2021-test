package controllers

import (
	"github.com/egorkartashov/xsolla-school-backend-2021-test/utils"
	"net/http"
)

var GetPing = func(writer http.ResponseWriter, request *http.Request) {
	response := map[string]interface{}{"message": "GET /PING"}
	utils.RespondJson(writer, http.StatusOK, response)
}
