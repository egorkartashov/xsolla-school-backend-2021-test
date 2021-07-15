package repos

import (
	"github.com/egorkartashov/xsolla-school-backend-2021-test/database/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductsRepo struct {
	db *gorm.DB
}

func NewProductsRepo(db *gorm.DB) *ProductsRepo {
	return &ProductsRepo{
		db: db,
	}
}

func (repo *ProductsRepo) GetProductOrNil(id uuid.UUID) *models.Product {
	// TODO lock or smth
	var product models.Product
	if err := repo.db.First(&product, id).Error; err != nil {
		return nil
	}
	return &product
}

func (repo *ProductsRepo) CreateProduct(product *models.Product) (*models.Product, error) {
	product.ID = uuid.New()
	if err := repo.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *ProductsRepo) UpdateProduct(product *models.Product) (*models.Product, error) {
	currentProduct := models.Product{}
	if err := repo.db.Where("id = ?", product.ID).First(&currentProduct).Error; err != nil {
		return nil, err
	}

	if err := repo.db.Model(currentProduct).Updates(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}
