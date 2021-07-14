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

type result struct {
	product *models.Product
	err     error
}

func NewProductsService(productsRepo *repos.ProductsRepo) *ProductsService {
	return &ProductsService{
		productsRepo: productsRepo,
	}
}

func (service *ProductsService) GetProduct(id uuid.UUID) (*dto.ProductDto, bool) {
	result := make(chan *models.Product)
	go func() {
		result <- service.productsRepo.GetProductOrNil(id)
	}()

	product := <-result
	if product == nil {
		return nil, false
	}

	productDto := mapModelToDto(product)
	return productDto, true
}

func (service *ProductsService) CreateProduct(productDto *dto.ProductDto) (*dto.ProductDto, error) {
	product := mapDtoToModel(productDto)
	resultChan := make(chan result)
	go func() {
		createdProduct, err := service.productsRepo.CreateProduct(product)
		resultChan <- result{product: createdProduct, err: err}
	}()

	result := <-resultChan
	if result.err != nil {
		return nil, result.err
	}

	createdProductDto := mapModelToDto(result.product)
	return createdProductDto, nil
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
