package repository

import (
	"gorm.io/gorm"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/model"
)

type OrderDetailRepository struct {
	Repository[model.OrderDetail]
}

func NewOrderDetailRepository(db *gorm.DB) *OrderDetailRepository {
	return &OrderDetailRepository{
		Repository: Repository[model.OrderDetail]{
			DB: db,
		},
	}
}

func (r *OrderDetailRepository) BulkCreate(tx *gorm.DB, data []*model.OrderDetail) error {
	db := r.DB
	if tx != nil {
		db = tx
	}
	return db.Create(data).Error
}
