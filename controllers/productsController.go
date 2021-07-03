package controllers

import (
	"github.com/egorkartashov/xsolla-school-backend-2021-test/services"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type ProductsController struct {
	productsService *services.ProductsService
}

func NewProductsController(productsService *services.ProductsService) *ProductsController {
	return &ProductsController{
		productsService: productsService,
	}
}

func (controller *ProductsController) GetProduct(writer http.ResponseWriter, request *http.Request) {
	productIdStr, found := mux.Vars(request)["id"]
	if found == false {
		respondBadId(writer)
		return
	}

	productId, err := uuid.Parse(productIdStr)
	if err != nil {
		respondBadId(writer)
		return
	}

	product, found := controller.productsService.GetProduct(productId)
	if found {
		utils.RespondJson(writer, http.StatusOK, product)
	} else {
		utils.RespondJson(writer, http.StatusNotFound, nil)
	}
}

func respondBadId(writer http.ResponseWriter) {
	errorResponse := make(map[string]interface{})
	errorResponse["message"] = "Could not parse product ID"
	utils.RespondJson(writer, http.StatusBadRequest, errorResponse)
}
