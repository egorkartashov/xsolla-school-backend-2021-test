package services

import (
	"errors"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/database/models"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/dto"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/repos"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductsService struct {
	productsRepo *repos.ProductsRepo
}

type singleProductResult struct {
	product *models.Product
	err     error
}

type productsListResult struct {
	productsList *[]models.Product
	err          error
}

func NewProductsService(productsRepo *repos.ProductsRepo) *ProductsService {
	return &ProductsService{
		productsRepo: productsRepo,
	}
}

func (service *ProductsService) GetProducts() (*[]dto.ProductDto, RequestResult) {
	repoResultChan := make(chan productsListResult)
	go func() {
		products, err := service.productsRepo.GetProducts()
		repoResultChan <- productsListResult{productsList: products, err: err}
	}()

	repoResult := <-repoResultChan
	requestResult := createRequestResult(repoResult.err)
	if requestResult.Status != Success {
		return nil, requestResult
	}

	var productsDtoList []dto.ProductDto
	for _, product := range *repoResult.productsList {
		productDto := mapModelToDto(&product)
		productsDtoList = append(productsDtoList, *productDto)
	}

	return &productsDtoList, requestResult
}

func (service *ProductsService) GetProduct(id uuid.UUID) (*dto.ProductDto, RequestResult) {
	resultChan := make(chan singleProductResult)
	go func() {
		product, err := service.productsRepo.GetProduct(id)
		resultChan <- singleProductResult{product: product, err: err}
	}()

	repoResult := <-resultChan

	requestResult := createRequestResult(repoResult.err)
	if requestResult.Status != Success {
		return nil, requestResult
	}

	productDto := mapModelToDto(repoResult.product)
	return productDto, requestResult
}

func (service *ProductsService) GetProductBySku(sku string) (*dto.ProductDto, RequestResult) {
	resultChan := make(chan singleProductResult)
	go func() {
		product, err := service.productsRepo.GetProductBySku(sku)
		resultChan <- singleProductResult{product: product, err: err}
	}()

	repoResult := <-resultChan

	requestResult := createRequestResult(repoResult.err)
	if requestResult.Status != Success {
		return nil, requestResult
	}

	productDto := mapModelToDto(repoResult.product)
	return productDto, requestResult
}

func (service *ProductsService) CreateProduct(productDto *dto.ProductDto) (*dto.ProductDto, RequestResult) {
	product := mapDtoToModel(productDto)
	resultChan := make(chan singleProductResult)
	go func() {
		createdProduct, err := service.productsRepo.CreateProduct(product)
		resultChan <- singleProductResult{product: createdProduct, err: err}
	}()

	result := <-resultChan

	requestResult := createRequestResult(result.err)
	if requestResult.Status != Success {
		return nil, requestResult
	}

	createdProductDto := mapModelToDto(result.product)
	return createdProductDto, requestResult
}

func (service *ProductsService) UpdateProduct(productDto *dto.ProductDto) (*dto.ProductDto, RequestResult) {
	_, err := service.productsRepo.GetProduct(*productDto.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		requestResult := RequestResult{Status: NotFound}
		return nil, requestResult
	}

	product := mapDtoToModel(productDto)
	resultChan := make(chan singleProductResult)
	go func() {
		createdProduct, err := service.productsRepo.UpdateProduct(product)
		resultChan <- singleProductResult{product: createdProduct, err: err}
	}()

	result := <-resultChan

	requestResult := createRequestResult(result.err)
	if requestResult.Status != Success {
		return nil, requestResult
	}

	updatedProductDto := mapModelToDto(result.product)
	return updatedProductDto, requestResult
}

func (service *ProductsService) UpdateProductBySku(productDto *dto.ProductDto) (*dto.ProductDto, RequestResult) {
	product := mapDtoToModel(productDto)
	resultChan := make(chan singleProductResult)
	go func() {
		updatedProduct, err := service.productsRepo.UpdateProductBySku(product)
		resultChan <- singleProductResult{product: updatedProduct, err: err}
	}()

	result := <-resultChan

	requestResult := createRequestResult(result.err)
	if requestResult.Status != Success {
		return nil, requestResult
	}

	updatedProductDto := mapModelToDto(result.product)
	return updatedProductDto, requestResult
}

func (service *ProductsService) DeleteProduct(productId uuid.UUID) RequestResult {
	errorChan := make(chan error)
	go func() {
		err := service.productsRepo.DeleteProduct(productId)
		errorChan <- err
	}()

	err := <-errorChan
	requestResult := createRequestResult(err)

	return requestResult
}

func (service *ProductsService) DeleteProductBySku(sku string) RequestResult {
	errorChan := make(chan error)
	go func() {
		err := service.productsRepo.DeleteProductBySku(sku)
		errorChan <- err
	}()

	err := <-errorChan
	requestResult := createRequestResult(err)

	return requestResult
}

func mapModelToDto(product *models.Product) *dto.ProductDto {
	return &dto.ProductDto{
		Id:           &product.ID,
		Sku:          product.Sku,
		Name:         product.Name,
		Type:         product.Type,
		PriceInCents: product.PriceInCents,
	}
}

func mapDtoToModel(productDto *dto.ProductDto) *models.Product {
	var id uuid.UUID
	if productDto.Id == nil {
		id, _ = uuid.NewUUID()
	} else {
		id = *productDto.Id
	}

	return &models.Product{
		ID:           id,
		Sku:          productDto.Sku,
		Name:         productDto.Name,
		Type:         productDto.Type,
		PriceInCents: productDto.PriceInCents,
	}
}

func createRequestResult(err error) RequestResult {
	requestResult := RequestResult{}
	if err == nil {
		requestResult.Status = Success
		return requestResult
	} else {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			requestResult.Status = NotFound
		} else {
			requestResult.Status = Error
			requestResult.Error = err
		}
		return requestResult
	}
}
