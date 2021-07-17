package services

import (
	"github.com/egorkartashov/xsolla-school-backend-2021-test/database/models"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/dto"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/repos"
	"github.com/google/uuid"
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

func (service *ProductsService) GetProducts() (*[]dto.ProductDto, error) {
	resultChan := make(chan productsListResult)
	go func() {
		products, err := service.productsRepo.GetProducts()
		resultChan <- productsListResult{productsList: products, err: err}
	}()

	result := <-resultChan
	if result.err != nil {
		return nil, result.err
	}

	var productsDtoList []dto.ProductDto
	for _, product := range *result.productsList {
		productDto := mapModelToDto(&product)
		productsDtoList = append(productsDtoList, *productDto)
	}

	return &productsDtoList, nil
}

func (service *ProductsService) GetProduct(id uuid.UUID) (*dto.ProductDto, error) {
	resultChan := make(chan singleProductResult)
	go func() {
		product, err := service.productsRepo.GetProduct(id)
		resultChan <- singleProductResult{product: product, err: err}
	}()

	result := <-resultChan
	if result.err != nil {
		return nil, result.err
	}

	productDto := mapModelToDto(result.product)
	return productDto, nil
}

func (service *ProductsService) GetProductBySku(sku string) (*dto.ProductDto, error) {
	resultChan := make(chan singleProductResult)
	go func() {
		product, err := service.productsRepo.GetProductBySku(sku)
		resultChan <- singleProductResult{product: product, err: err}
	}()

	result := <-resultChan
	if result.err != nil {
		return nil, result.err
	}

	productDto := mapModelToDto(result.product)
	return productDto, nil
}

func (service *ProductsService) CreateProduct(productDto *dto.ProductDto) (*dto.ProductDto, error) {
	product := mapDtoToModel(productDto)
	resultChan := make(chan singleProductResult)
	go func() {
		createdProduct, err := service.productsRepo.CreateProduct(product)
		resultChan <- singleProductResult{product: createdProduct, err: err}
	}()

	result := <-resultChan
	if result.err != nil {
		return nil, result.err
	}

	createdProductDto := mapModelToDto(result.product)
	return createdProductDto, nil
}

func (service *ProductsService) UpdateProduct(productDto *dto.ProductDto) (*dto.ProductDto, error) {
	product := mapDtoToModel(productDto)
	resultChan := make(chan singleProductResult)
	go func() {
		createdProduct, err := service.productsRepo.UpdateProduct(product)
		resultChan <- singleProductResult{product: createdProduct, err: err}
	}()

	result := <-resultChan
	if result.err != nil {
		return nil, result.err
	}

	updatedProductDto := mapModelToDto(result.product)
	return updatedProductDto, nil
}

func (service *ProductsService) UpdateProductBySku(productDto *dto.ProductDto) (*dto.ProductDto, error) {
	product := mapDtoToModel(productDto)
	resultChan := make(chan singleProductResult)
	go func() {
		updatedProduct, err := service.productsRepo.UpdateProductBySku(product)
		resultChan <- singleProductResult{product: updatedProduct, err: err}
	}()

	result := <-resultChan
	if result.err != nil {
		return nil, result.err
	}

	updatedProductDto := mapModelToDto(result.product)
	return updatedProductDto, nil
}

func (service *ProductsService) DeleteProduct(productId uuid.UUID) error {
	errorChan := make(chan error)
	go func() {
		err := service.productsRepo.DeleteProduct(productId)
		errorChan <- err
	}()

	err := <-errorChan
	return err
}

func (service *ProductsService) DeleteProductBySku(sku string) error {
	errorChan := make(chan error)
	go func() {
		err := service.productsRepo.DeleteProductBySku(sku)
		errorChan <- err
	}()

	err := <-errorChan
	return err
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
