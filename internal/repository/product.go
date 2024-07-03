package repository

import (
	"gorm.io/gorm"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/domain"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/model"
)

type ProductRepository struct {
	Repository[model.Product]
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		Repository: Repository[model.Product]{
			DB: db,
		},
	}
}

func (r *ProductRepository) Find(request *domain.GetProductsRequest) ([]model.Product, uint16, error) {
	var (
		total  int64
		result []model.Product
	)

	filter := func(db *gorm.DB) *gorm.DB {
		if request.CategoryID > 0 {
			db = db.Where("category_id = ?", request.CategoryID)
		}
		return db
	}

	if err := r.DB.Preload("Category").Scopes(filter).Offset(int((request.Page - 1) * request.Size)).Limit(int(request.Size)).Find(&result).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Preload("Category").Model(new(model.Product)).Scopes(filter).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return result, uint16(total), nil
}

func (r *ProductRepository) FindByID(productID uint64) (*model.Product, error) {
	var product model.Product
	err := r.DB.Preload("Category").First(&product, productID).Error
	return &product, err
}
