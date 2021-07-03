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
