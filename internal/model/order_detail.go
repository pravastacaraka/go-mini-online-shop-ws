package model

import "time"

type OrderDetail struct {
	ID             uint64    `gorm:"column:id;primaryKey"`
	OrderID        uint64    `gorm:"column:order_id"`
	ProductID      uint64    `gorm:"column:product_id"`
	ProductName    string    `gorm:"column:product_name"`
	Quantity       uint16    `gorm:"column:quantity"`
	Price          uint32    `gorm:"column:price"`
	SubTotalPrice  uint32    `gorm:"column:subtotal_price"`
	Weight         float32   `gorm:"column:weight"`
	SubTotalWeight float32   `gorm:"column:subtotal_weight"`
	CategoryID     uint64    `gorm:"column:category_id"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (o *OrderDetail) TableName() string {
	return "order_detail"
}
