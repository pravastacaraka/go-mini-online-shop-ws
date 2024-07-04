package repository

import (
	"gorm.io/gorm"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/domain"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/model"
)

type OrderRepository struct {
	Repository[model.Order]
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		Repository: Repository[model.Order]{
			DB: db,
		},
	}
}

func (r *OrderRepository) FindByPaymentID(id any) (*model.Order, error) {
	var data *model.Order
	err := r.DB.Preload("OrderDetails").Where("payment_id = ?", id).First(&data).Error
	return data, err
}

func (r *OrderRepository) Find(request *domain.GetOrdersRequest) ([]model.Order, uint16, error) {
	var (
		total  int64
		result []model.Order
	)

	filter := func(db *gorm.DB) *gorm.DB {
		if request.UserID > 0 {
			db = db.Where("user_id = ?", request.UserID)
		}
		if request.Invoice != "" {
			db = db.Where("invoice LIKE ?", "%"+request.Invoice+"%")
		}
		if !request.StartDate.IsZero() {
			db = db.Where("created_at >= ?", request.StartDate)
		}
		if !request.EndDate.IsZero() {
			db = db.Where("created_at <= ?", request.EndDate)
		}
		return db
	}

	if err := r.DB.Preload("User").
		Preload("Address").
		Preload("Payment").
		Preload("OrderDetails").
		Scopes(filter).
		Offset(int((request.Page - 1) * request.Size)).
		Limit(int(request.Size)).
		Find(&result).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.Model(&model.Order{}).Scopes(filter).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return result, uint16(total), nil
}
