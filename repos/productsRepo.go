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

func (repo *ProductsRepo) GetProducts(offset, limit int) (*[]models.Product, error) {
	var products []models.Product
	if err := repo.db.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
		return nil, err
	}

	return &products, nil
}

func (repo *ProductsRepo) GetProduct(id uuid.UUID) (*models.Product, error) {
	var product models.Product
	if err := repo.db.First(&product, id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo *ProductsRepo) GetProductBySku(sku string) (*models.Product, error) {
	var product models.Product
	if err := repo.db.Where("sku = ?", sku).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
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

func (repo *ProductsRepo) UpdateProductBySku(product *models.Product) (*models.Product, error) {
	currentProduct := models.Product{}
	if err := repo.db.Where("sku = ?", product.Sku).First(&currentProduct).Error; err != nil {
		return nil, err
	}

	product.ID = currentProduct.ID
	if err := repo.db.Model(currentProduct).Updates(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (repo *ProductsRepo) DeleteProduct(productId uuid.UUID) error {
	err := repo.db.Unscoped().Delete(&models.Product{}, productId).Error

	return err
}

func (repo *ProductsRepo) DeleteProductBySku(sku string) error {
	err := repo.db.Unscoped().Where("sku = ?", sku).Delete(&models.Product{}).Error

	return err
}

func (repo *ProductsRepo) GetAllTypes() ([]string, error) {
	var products []models.Product
	var types = make([]string, 0)
	if err := repo.db.Model(&products).Distinct().Pluck("type", &types).Error; err != nil {
		return types, err
	}

	return types, nil
}
