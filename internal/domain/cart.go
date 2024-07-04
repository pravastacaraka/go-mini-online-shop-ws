package domain

type AddToCartRequest struct {
	UserID    uint64 `json:"user_id" validate:"required,min=1"`
	ProductID uint64 `json:"product_id" validate:"required,min=1"`
	Quantity  uint16 `json:"quantity" validate:"required,min=1"`
}

type DeleteCartRequest struct {
	CartID uint64 `json:"cart_id" validate:"required,min=1"`
}

type UpdateCartDetailRequest struct {
	CartDetailID uint64 `json:"cart_detail_id" validate:"required,min=1"`
	Quantity     uint16 `json:"quantity" validate:"required,min=1"`
}

type DeleteCartDetailRequest struct {
	CartDetailID uint64 `json:"cart_detail_id" validate:"required,min=1"`
}

type CommonCartResponse struct {
	Message string `json:"message"`
}

type GetCartListRequest struct {
	UserID uint64 `json:"user_id" validate:"required,min=1"`
}

type GetCartListResponse struct {
	CartID               uint64      `json:"cart_id"`
	AvailableProducts    []*CartList `json:"available_products"`
	NonAvailableProducts []*CartList `json:"non_available_products"`
}

type CartList struct {
	ProductID    uint64 `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductPrice uint32 `json:"product_price"`
	Quantity     uint16 `json:"quantity"`
}

type CheckoutRequest struct {
	CartID uint64 `json:"cart_id" validate:"required,min=1"`
	UserID uint64 `json:"user_id" validate:"required,min=1"`
}

type CheckoutResponse struct {
	CartID       uint64         `json:"cart_id"`
	Customer     Customer       `json:"customer"`
	Payment      Payment        `json:"payment"`
	Shipment     Shipment       `json:"shipment"`
	OrderDetails []*OrderDetail `json:"order_details"`
}
