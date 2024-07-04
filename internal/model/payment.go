package model

import "time"

type Payment struct {
	ID          uint64        `gorm:"column:id;primaryKey"`
	CartID      uint64        `gorm:"column:cart_id;primaryKey"`
	Amount      uint32        `gorm:"column:total_payment"`
	GatewayName string        `gorm:"column:gateway_name"`
	Status      PaymentStatus `gorm:"column:status;type:integer"`
	CreatedAt   time.Time     `gorm:"column:created_at"`
	UpdatedAt   time.Time     `gorm:"column:updated_at"`

	// Just for dummy payment, will be generated after order creation process.
	// This code is used in the body of "/api/v1/order/pay/:paymentId".
	PaymentCode string `gorm:"column:payment_code"`
}

func (p *Payment) TableName() string {
	return "payment"
}

type PaymentStatus int8

const (
	PaymentStatusPending           PaymentStatus = -1
	PaymentStatusCanceled          PaymentStatus = 0
	PaymentStatusWaitingThirdParty PaymentStatus = 1
	PaymentStatusSucceed           PaymentStatus = 2
)
