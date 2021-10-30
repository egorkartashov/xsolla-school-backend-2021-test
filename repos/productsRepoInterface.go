package repos

import (
	"github.com/egorkartashov/xsolla-school-backend-2021-test/database/models"
	"github.com/egorkartashov/xsolla-school-backend-2021-test/filters"
	"github.com/google/uuid"
)

type ProductsRepoInterface interface {
	GetProducts(filters []filters.FilterPair, offset, limit int) (*[]models.Product, error)
	GetProduct(id uuid.UUID) (*models.Product, error)
	GetProductBySku(sku string) (*models.Product, error)
	CreateProduct(product *models.Product) (*models.Product, error)
	UpdateProduct(product *models.Product) (*models.Product, error)
	UpdateProductBySku(product *models.Product) (*models.Product, error)
	DeleteProduct(productId uuid.UUID) error
	DeleteProductBySku(sku string) error
	GetAllTypes() ([]string, error)
}
