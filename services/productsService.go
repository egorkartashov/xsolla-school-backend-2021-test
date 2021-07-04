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

func (service *ProductsService) CreateProduct(productDto *dto.ProductDto) (*dto.ProductDto, bool) {
	product := mapDtoToModel(productDto)
	errChan := make(chan error)
	productChan := make(chan *models.Product)
	go func() {
		createdProduct, err := service.productsRepo.CreateProduct(product)
		productChan <- createdProduct
		errChan <- err
	}()

	err := <-errChan
	product = <-productChan
	if err != nil {
		return nil, false
	}

	createdProductDto := mapModelToDto(product)
	return createdProductDto, true
}

func mapModelToDto(product *models.Product) *dto.ProductDto {
	return &dto.ProductDto{
		Id:           product.ID,
		Sku:          product.Sku,
		Name:         product.Name,
		Type:         product.Type,
		PriceInCents: product.PriceInCents,
	}
}

func mapDtoToModel(productDto *dto.ProductDto) *models.Product {
	return &models.Product{
		ID:           productDto.Id,
		Sku:          productDto.Sku,
		Name:         productDto.Name,
		Type:         productDto.Type,
		PriceInCents: productDto.PriceInCents,
	}
}
