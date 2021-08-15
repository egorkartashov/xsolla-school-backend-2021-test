package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/dto"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/filters"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/services"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/utils"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
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

const (
	offsetParamKey      = "offset"
	limitParamKey       = "limit"
	productTypeParamKey = "type"
	minPriceParamKey    = "min_price"
	maxPriceParamKey    = "max_price"
	defaultOffsetValue  = 0
	defaultLimitValue   = 100
)

func (controller *ProductsController) GetProducts(writer http.ResponseWriter, request *http.Request) {
	var offset, limit int
	var ok bool

	if offset, ok = utils.TryParseIntQueryParameterOrDefault(request, offsetParamKey, defaultOffsetValue); !ok {
		utils.RespondErrorJson(writer, http.StatusBadRequest,
			fmt.Sprintf("Could not parse \"%s\" query parameter", offsetParamKey))
		return
	}

	if limit, ok = utils.TryParseIntQueryParameterOrDefault(request, limitParamKey, defaultLimitValue); !ok {
		utils.RespondErrorJson(writer, http.StatusBadRequest,
			fmt.Sprintf("Could not parse \"%s\" query parameter", limitParamKey))
		return
	}

	var minPrice, maxPrice *int
	var productType *string
	var err error

	var productTypeStr = request.URL.Query().Get(productTypeParamKey)
	if productTypeStr != "" {
		productType = &productTypeStr
	}

	if minPrice, err = utils.TryParseIntQueryParameterOrNil(request, minPriceParamKey); err != nil {
		utils.RespondErrorJson(writer, http.StatusBadRequest,
			fmt.Sprintf("Could not parse \"%s\" query parameter", minPriceParamKey))
		return
	}

	if maxPrice, err = utils.TryParseIntQueryParameterOrNil(request, maxPriceParamKey); err != nil {
		utils.RespondErrorJson(writer, http.StatusBadRequest,
			fmt.Sprintf("Could not parse \"%s\" query parameter", maxPriceParamKey))
		return
	}

	productsFilters := filters.BuildProductsFilters(productType, minPrice, maxPrice)

	products, requestResult := controller.productsService.GetProducts(&productsFilters, offset, limit)
	if requestResult.Status == services.Success {
		size := len(products)
		pagination := dto.CreatePaginationWithLinks(request.URL.Path, offset, limit, size)

		paginatedProductsDto := dto.PaginatedProductsDto{
			Pagination: pagination,
			Data:       products,
		}
		utils.RespondJson(writer, http.StatusOK, paginatedProductsDto)
	} else if requestResult.Status == services.NotFound {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		handleUnknownError(writer, request, requestResult)
	}
}

func (controller *ProductsController) PostProduct(writer http.ResponseWriter, request *http.Request) {
	productDto, ok := parseAndValidateProductDto(writer, request)
	if !ok {
		return
	}

	createdProductDto, requestResult := controller.productsService.CreateProduct(productDto)
	if requestResult.Status == services.Success {
		utils.RespondJson(writer, http.StatusCreated, createdProductDto)
	} else {
		handleUnknownError(writer, request, requestResult)
	}
}

func (controller *ProductsController) GetProductBySku(writer http.ResponseWriter, request *http.Request) {
	productSku, found := mux.Vars(request)["sku"]
	if !found {
		utils.RespondErrorJson(writer, http.StatusBadRequest, "Could not find product SKU")
		return
	}

	product, requestResult := controller.productsService.GetProductBySku(productSku)
	if requestResult.Status == services.Success {
		utils.RespondJson(writer, http.StatusOK, product)
	} else if requestResult.Status == services.NotFound {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		handleUnknownError(writer, request, requestResult)
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
	updatedProductDto, requestResult := controller.productsService.UpdateProductBySku(productDto)
	if requestResult.Status == services.Success {
		utils.RespondJson(writer, http.StatusOK, updatedProductDto)
	} else if requestResult.Status == services.NotFound {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		handleUnknownError(writer, request, requestResult)
	}
}

func (controller *ProductsController) DeleteProductBySku(writer http.ResponseWriter, request *http.Request) {
	productSku, found := mux.Vars(request)["sku"]
	if !found {
		utils.RespondErrorJson(writer, http.StatusBadRequest, "Could not find product SKU")
		return
	}

	requestResult := controller.productsService.DeleteProductBySku(productSku)

	if requestResult.Status == services.Success {
		writer.WriteHeader(http.StatusNoContent)
	} else if requestResult.Status == services.NotFound {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		handleUnknownError(writer, request, requestResult)
	}
}

func (controller *ProductsController) GetProduct(writer http.ResponseWriter, request *http.Request) {
	productId, parsed := tryParseProductId(writer, request)
	if !parsed {
		return
	}

	product, requestResult := controller.productsService.GetProduct(productId)
	if requestResult.Status == services.Success {
		utils.RespondJson(writer, http.StatusOK, product)
	} else if requestResult.Status == services.NotFound {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		handleUnknownError(writer, request, requestResult)
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
	updatedProductDto, requestResult := controller.productsService.UpdateProduct(productDto)
	if requestResult.Status == services.Success {
		utils.RespondJson(writer, http.StatusOK, updatedProductDto)
	} else if requestResult.Status == services.NotFound {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		handleUnknownError(writer, request, requestResult)
	}
}

func (controller *ProductsController) DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	productId, parsed := tryParseProductId(writer, request)
	if !parsed {
		return
	}

	requestResult := controller.productsService.DeleteProduct(productId)
	if requestResult.Status == services.Success {
		writer.WriteHeader(http.StatusNoContent)
	} else if requestResult.Status == services.NotFound {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		handleUnknownError(writer, request, requestResult)
	}
}

func (controller *ProductsController) GetAllTypes(writer http.ResponseWriter, request *http.Request) {
	types, requestResult := controller.productsService.GetAllTypes()
	if requestResult.Status == services.Success {
		utils.RespondJson(writer, http.StatusOK, types)
	} else {
		handleUnknownError(writer, request, requestResult)
	}
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

func handleUnknownError(writer http.ResponseWriter, request *http.Request, requestResult services.RequestResult) {
	route, _ := mux.CurrentRoute(request).GetPathTemplate()
	log.Printf("Error in %s: status = %s, error = %s", route, requestResult.Status, requestResult.Error)
	writer.WriteHeader(http.StatusInternalServerError)
}
