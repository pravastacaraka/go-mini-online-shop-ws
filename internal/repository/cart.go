package repository

import (
	"gorm.io/gorm"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/model"
)

type CartRepository struct {
	Repository[model.Cart]
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		Repository: Repository[model.Cart]{
			DB: db,
		},
	}
}

func (r *CartRepository) CountByUserID(id any) (int64, error) {
	var total int64
	err := r.DB.Model(&model.Cart{}).Where("user_id = ?", id).Count(&total).Error
	return total, err
}

func (r *CartRepository) FindByUserID(id any) (*model.Cart, error) {
	var cart *model.Cart

	err := r.DB.
		Preload("CartDetails", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "CartID", "Quantity", "ProductID")
		}).
		Preload("CartDetails.Product", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID", "Name", "Price", "Stock")
		}).
		Where("user_id = ?", id).
		First(&cart).Error

	return cart, err
}
