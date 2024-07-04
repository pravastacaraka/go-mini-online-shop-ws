package repository

import (
	"gorm.io/gorm"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/model"
)

type PaymentRepository struct {
	Repository[model.Payment]
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		Repository: Repository[model.Payment]{
			DB: db,
		},
	}
}

func (r *PaymentRepository) CountByCartID(id any) (int64, error) {
	var total int64
	err := r.DB.Model(&model.Payment{}).Where("cart_id = ?", id).Count(&total).Error
	return total, err
}
