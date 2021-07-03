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

	productDto := MapModelToDto(product)
	return productDto, true
}

func MapModelToDto(product *models.Product) *dto.ProductDto {
	return &dto.ProductDto{
		Id:           product.Id,
		Sku:          product.Sku,
		Name:         product.Name,
		Type:         product.Type,
		PriceInCents: product.PriceInCents,
	}
}
