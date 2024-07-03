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
	AvailableProducts    []*CartList `json:"available_products"`
	NonAvailableProducts []*CartList `json:"non_available_products"`
}

type CartList struct {
	ID           uint64 `json:"id"`
	ProductID    uint64 `json:"product_id"`
	ProductName  string `json:"product_name"`
	ProductPrice uint32 `json:"product_price"`
	Quantity     uint16 `json:"quantity"`
}
