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

func (controller *ProductsController) GetProducts(writer http.ResponseWriter, _ *http.Request) {
	// TODO pagination (offset, limit)

	products, err := controller.productsService.GetProducts()
	if err == nil {
		utils.RespondJson(writer, http.StatusOK, products)
	} else {
		utils.RespondErrorJson(writer, http.StatusNotFound, fmt.Sprintf("Could not get products: %s", err))
	}
}

func (controller *ProductsController) PostProduct(writer http.ResponseWriter, request *http.Request) {
	productDto, ok := parseAndValidateProductDto(writer, request)
	if !ok {
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

func (controller *ProductsController) GetProductBySku(writer http.ResponseWriter, request *http.Request) {
	productSku, found := mux.Vars(request)["sku"]
	if !found {
		utils.RespondErrorJson(writer, http.StatusBadRequest, "Could not find product SKU")
		return
	}

	product, err := controller.productsService.GetProductBySku(productSku)
	if err == nil {
		utils.RespondJson(writer, http.StatusOK, product)
	} else {
		utils.RespondJson(writer, http.StatusNotFound, fmt.Sprintf("Could not get product: %s", err))
	}
}

func (controller *ProductsController) PutProductBySku(writer http.ResponseWriter, request *http.Request) {
	productSku, found := mux.Vars(request)["sku"]
	if !found {
		utils.RespondErrorJson(writer, http.StatusBadRequest, "Could not find product SKU")
		return
	}

	productDto, ok := parseAndValidateProductDto(writer, request)
	if !ok {
		return
	}

	productDto.Sku = productSku
	updatedProductDto, err := controller.productsService.UpdateProductBySku(productDto)
	if err == nil {
		utils.RespondJson(writer, http.StatusOK, updatedProductDto)
	} else {
		utils.RespondErrorJson(writer, http.StatusInternalServerError,
			fmt.Sprintf("Could not update product: %s", err))
	}
}

func (controller *ProductsController) DeleteProductBySku(writer http.ResponseWriter, request *http.Request) {
	productSku, found := mux.Vars(request)["sku"]
	if !found {
		utils.RespondErrorJson(writer, http.StatusBadRequest, "Could not find product SKU")
		return
	}

	err := controller.productsService.DeleteProductBySku(productSku)
	if err == nil {
		writer.WriteHeader(http.StatusOK)
	} else {
		utils.RespondErrorJson(writer, http.StatusInternalServerError,
			fmt.Sprintf("Could not delete product: %s", err))
	}
}

func (controller *ProductsController) GetProduct(writer http.ResponseWriter, request *http.Request) {
	productId, parsed := tryParseProductId(writer, request)
	if !parsed {
		return
	}

	product, err := controller.productsService.GetProduct(productId)
	if err == nil {
		utils.RespondJson(writer, http.StatusOK, product)
	} else {
		utils.RespondJson(writer, http.StatusNotFound, fmt.Sprintf("Could not get product: %s", err))
	}
}

func (controller *ProductsController) PutProduct(writer http.ResponseWriter, request *http.Request) {
	productId, parsed := tryParseProductId(writer, request)
	if !parsed {
		return
	}

	productDto, ok := parseAndValidateProductDto(writer, request)
	if !ok {
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

func (controller *ProductsController) DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	productId, parsed := tryParseProductId(writer, request)
	if !parsed {
		return
	}

	err := controller.productsService.DeleteProduct(productId)
	if err == nil {
		writer.WriteHeader(http.StatusOK)
	} else {
		utils.RespondErrorJson(writer, http.StatusInternalServerError,
			fmt.Sprintf("Could not delete product: %s", err))
	}
}

func (controller *ProductsController) GetAllTypes(writer http.ResponseWriter, request *http.Request) {
	utils.RespondErrorJson(writer, http.StatusInternalServerError, "Endpoint not implemented yet")
}

func tryParseProductId(writer http.ResponseWriter, request *http.Request) (uuid.UUID, bool) {
	productIdStr, found := mux.Vars(request)["id"]
	if !found {
		utils.RespondErrorJson(writer, http.StatusBadRequest, "Could not find product ID")
		return uuid.UUID{}, false
	}

	productId, err := uuid.Parse(productIdStr)
	if err != nil {
		utils.RespondErrorJson(writer, http.StatusBadRequest, "Could not parse product ID")
		return uuid.UUID{}, false
	}

	return productId, true
}

func parseAndValidateProductDto(writer http.ResponseWriter, request *http.Request) (*dto.ProductDto, bool) {
	productDto := &dto.ProductDto{}
	if err := json.NewDecoder(request.Body).Decode(productDto); err != nil {
		utils.RespondErrorJson(writer, http.StatusBadRequest, "Could not parse request body")
		return nil, false
	}

	v := validator.New()
	err := v.Struct(productDto)
	if err != nil {
		utils.RespondValidationErrors(writer, err.(validator.ValidationErrors))
		return nil, false
	}

	return productDto, true
}
