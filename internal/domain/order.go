package domain

import "time"

type CreateOrderRequest struct {
	CartID       uint64         `json:"cart_id" validate:"required,min=1"`
	Payment      Payment        `json:"payment" validate:"required"`
	Customer     Customer       `json:"customer" validate:"required"`
	Shipment     Shipment       `json:"shipment" validate:"required"`
	OrderDetails []*OrderDetail `json:"order_details" validate:"required"`
}

type CreateOrderResponse struct {
	Invoice     string `json:"invoice,omitempty"`
	PaymentID   uint64 `json:"payment_id,omitempty"`
	PaymentCode string `json:"payment_code,omitempty"`
	Message     string `json:"message,omitempty"`
}

type OrderDetail struct {
	ProductID      uint64  `json:"product_id" validate:"required,min=1"`
	ProductName    string  `json:"product_name" validate:"required"`
	Quantity       uint16  `json:"quantity" validate:"required"`
	Price          uint32  `json:"price"`
	SubTotalPrice  uint32  `json:"subtotal_price"`
	Weight         float32 `json:"weight" validate:"required"`
	SubTotalWeight float32 `json:"subtotal_weight" validate:"required"`
	CategoryID     uint64  `json:"category_id" validate:"required"`
}

type Customer struct {
	UserID uint64 `json:"user_id" validate:"required,min=1"`
}

type Payment struct {
	TotalPrice   uint32 `json:"total_price"`
	TotalPayment uint32 `json:"total_payment"`
}

type Shipment struct {
	AddressID     uint64  `json:"address_id" validate:"required,min=1"`
	ShippingPrice uint32  `json:"shipping_price"`
	TotalWeight   float32 `json:"total_weight" validate:"required"`
}

type GetOrdersRequest struct {
	UserID    uint64    `json:"user_id" validate:"required,min=1"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Invoice   string    `json:"invoice"`
	Page      uint8     `json:"page" validate:"required,numeric,min=1"`
	Size      uint8     `json:"size" validate:"required,numeric,min=1"`
}

type GetOrderBuyerResponse struct {
	Invoice      string                               `json:"invoice"`
	Status       string                               `json:"status"`
	TotalPrice   string                               `json:"total_price"`
	OrderDetails []*GetOrderBuyerResponse_OrderDetail `json:"order_details"`
	Customer     GetOrderBuyerResponse_Customer       `json:"customer"`
	Shipment     GetOrderBuyerResponse_Shipment       `json:"shipment"`
	Payment      GetOrderBuyerResponse_Payment        `json:"payment"`
}

type GetOrderBuyerResponse_OrderDetail struct {
	ProductName string  `json:"product_name"`
	Quantity    uint16  `json:"quantity"`
	Price       string  `json:"price"`
	Weight      float32 `json:"weight"`
	CategoryID  uint64  `json:"category_id"`
}

type GetOrderBuyerResponse_Customer struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	PostalCode uint32 `json:"postal_code"`
}

type GetOrderBuyerResponse_Shipment struct {
	ShippingAgentName    string  `json:"shipping_agent_name"`
	ShippingAgentProduct string  `json:"shipping_agent_product"`
	ShippingPrice        string  `json:"shipping_price"`
	TotalWeight          float32 `json:"total_weight"`
}

type GetOrderBuyerResponse_Payment struct {
	GatewayName  string `json:"gateway_name"`
	TotalPayment string `json:"total_payment"`
}

type DoPaymentOrderRequest struct {
	UserID      uint64 `json:"user_id" validate:"required,min=1"`
	PaymentID   uint64 `json:"payment_id" validate:"required,min=1"`
	PaymentCode string `json:"payment_code" validate:"required"`
}

type DoPaymentOrderResponse struct {
	Invoice string `json:"invoice,omitempty"`
	Message string `json:"message,omitempty"`
}
