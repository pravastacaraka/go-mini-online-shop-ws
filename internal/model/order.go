package model

import "time"

type Order struct {
	ID            uint64         `gorm:"column:id;primaryKey"`
	PaymentID     uint64         `gorm:"column:payment_id"`
	UserID        uint64         `gorm:"column:user_id"`
	AddressID     uint64         `gorm:"column:address_id"`
	Invoice       string         `gorm:"column:invoice"`
	TotalPrice    uint32         `gorm:"column:total_price"`
	TotalWeight   float32        `gorm:"column:total_weight"`
	ShippingPrice uint32         `gorm:"column:shipping_price"`
	Status        OrderStatus    `gorm:"column:status;type:text"`
	CreatedAt     time.Time      `gorm:"column:created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at"`
	User          *User          `gorm:"foreignKey:user_id"`
	Address       *Address       `gorm:"foreignKey:address_id"`
	Payment       *Payment       `gorm:"foreignKey:payment_id"`
	OrderDetails  []*OrderDetail `gorm:"foreignKey:order_id"`
}

func (o *Order) TableName() string {
	return "order"
}

type OrderStatus uint16

const (
	OrderStatusCanceled  OrderStatus = 0
	OrderStatusCreated   OrderStatus = 100
	OrderStatusProcessed OrderStatus = 200
	OrderStatusShipped   OrderStatus = 300
	OrderStatusDelivered OrderStatus = 500
	OrderStatusFinished  OrderStatus = 600
)

var OrderStatusDesc = map[OrderStatus]string{
	OrderStatusCanceled:  "Dibatalkan",
	OrderStatusCreated:   "Menunggu Pembayaran",
	OrderStatusProcessed: "Diproses",
	OrderStatusShipped:   "Dikirim",
	OrderStatusDelivered: "Tiba di Tujuan",
	OrderStatusFinished:  "Selesai",
}
