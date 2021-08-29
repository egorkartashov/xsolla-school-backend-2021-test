package caching

import (
	"github.com/egorkartashov/xsolla-school-backend-2021-test/database/models"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/filters"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/repos"
	"github.com/google/uuid"
)

type ProductsRepoWithCache struct {
	repoInterface repos.ProductsRepoInterface
}

func (repo *ProductsRepoWithCache) GetProducts(filters []filters.FilterPair, offset, limit int) (*[]models.Product, error) {
	return repo.repoInterface.GetProducts(filters, offset, limit)
}

func (repo *ProductsRepoWithCache) GetProduct(id uuid.UUID) (*models.Product, error) {
	return repo.repoInterface.GetProduct(id)
}

func (repo *ProductsRepoWithCache) GetProductBySku(sku string) (*models.Product, error) {
	return repo.repoInterface.GetProductBySku(sku)
}

func (repo *ProductsRepoWithCache) CreateProduct(product *models.Product) (*models.Product, error) {
	return repo.repoInterface.CreateProduct(product)
}

func (repo *ProductsRepoWithCache) UpdateProduct(product *models.Product) (*models.Product, error) {
	return repo.repoInterface.UpdateProduct(product)
}

func (repo *ProductsRepoWithCache) UpdateProductBySku(product *models.Product) (*models.Product, error) {
	return repo.repoInterface.UpdateProductBySku(product)
}

func (repo *ProductsRepoWithCache) DeleteProduct(productId uuid.UUID) error {
	return repo.repoInterface.DeleteProduct(productId)
}

func (repo *ProductsRepoWithCache) DeleteProductBySku(sku string) error {
	return repo.repoInterface.DeleteProductBySku(sku)
}

func (repo *ProductsRepoWithCache) GetAllTypes() ([]string, error) {
	return repo.repoInterface.GetAllTypes()
}

func NewProductsRepoWithCaching(repoInterface repos.ProductsRepoInterface) *ProductsRepoWithCache {
	return &ProductsRepoWithCache{
		repoInterface: repoInterface,
	}
}
