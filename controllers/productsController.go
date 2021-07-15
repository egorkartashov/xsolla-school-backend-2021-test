package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/dto"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/services"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/utils"
	"github.com/go-playground/validator"
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
	if !found {
		utils.RespondErrorJson(writer, http.StatusBadRequest, "Could not find product ID")
		return
	}

	productId, err := uuid.Parse(productIdStr)
	if err != nil {
		utils.RespondErrorJson(writer, http.StatusBadRequest, "Could not parse product ID")
		return
	}

	product, found := controller.productsService.GetProduct(productId)
	if found {
		utils.RespondJson(writer, http.StatusOK, product)
	} else {
		utils.RespondJson(writer, http.StatusNotFound, nil)
	}
}

func (controller *ProductsController) PostProduct(writer http.ResponseWriter, request *http.Request) {
	productDto := &dto.ProductDto{}
	if err := json.NewDecoder(request.Body).Decode(productDto); err != nil {
		utils.RespondErrorJson(writer, http.StatusBadRequest, "Could not parse request body")
		return
	}

	v := validator.New()
	err := v.Struct(productDto)
	if err != nil {
		utils.RespondValidationErrors(writer, err.(validator.ValidationErrors))
		return
	}

	createdProductDto, err := controller.productsService.CreateProduct(productDto)
	if err == nil {
		utils.RespondJson(writer, http.StatusCreated, createdProductDto)
	} else {
		utils.RespondErrorJson(writer, http.StatusInternalServerError,
			fmt.Sprintf("Could not create product: %s", err))
	}
}

func (controller *ProductsController) PutProduct(writer http.ResponseWriter, request *http.Request) {
	productIdStr, found := mux.Vars(request)["id"]
	if !found {
		utils.RespondErrorJson(writer, http.StatusBadRequest, "Could not find product ID")
		return
	}

	productId, err := uuid.Parse(productIdStr)
	if err != nil {
		utils.RespondErrorJson(writer, http.StatusBadRequest, "Could not parse product ID")
		return
	}

	productDto := &dto.ProductDto{}
	if err := json.NewDecoder(request.Body).Decode(productDto); err != nil {
		utils.RespondErrorJson(writer, http.StatusBadRequest, "Could not parse request body")
		return
	}

	v := validator.New()
	err = v.Struct(productDto)
	if err != nil {
		utils.RespondValidationErrors(writer, err.(validator.ValidationErrors))
		return
	}

	productDto.Id = &productId
	updatedProductDto, err := controller.productsService.UpdateProduct(productDto)
	if err == nil {
		utils.RespondJson(writer, http.StatusOK, updatedProductDto)
	} else {
		utils.RespondErrorJson(writer, http.StatusInternalServerError,
			fmt.Sprintf("Could not update product: %s", err))
	}
}

func (controller *ProductsController) DeleteProductById(writer http.ResponseWriter, request *http.Request) {
	utils.RespondErrorJson(writer, http.StatusInternalServerError, "Endpoint not implemented yet")
}

func (controller *ProductsController) DeleteProductBySku(writer http.ResponseWriter, request *http.Request) {
	utils.RespondErrorJson(writer, http.StatusInternalServerError, "Endpoint not implemented yet")
}

func (controller *ProductsController) GetAllTypes(writer http.ResponseWriter, request *http.Request) {
	utils.RespondErrorJson(writer, http.StatusInternalServerError, "Endpoint not implemented yet")
}
