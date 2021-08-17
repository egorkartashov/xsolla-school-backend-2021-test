package controllers

import (
	"github.com/egorkartashov/xsolla-school-backend-2021-test/utils"
	"net/http"
)

type LoginController struct {
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (controller *LoginController) PostServerLogin(writer http.ResponseWriter, request *http.Request) {
	utils.RespondJson(writer, http.StatusOK, "server-login")
}

func (controller *LoginController) GetClientLogin(writer http.ResponseWriter, request *http.Request) {
	utils.RespondJson(writer, http.StatusOK, "client-login")
}
