package domain

type CommonProductResponse struct {
	ID uint64 `json:"id"`
}

type AddProductRequest struct {
	Name       string   `json:"name" validate:"required,max=150"`
	Desc       string   `json:"desc" validate:"required,max=1000"`
	SKU        string   `json:"sku"`
	Stock      uint16   `json:"stock" validate:"required,numeric,max=999"`
	Price      uint32   `json:"price" validate:"required,numeric,max=1000000000"`
	CategoryID uint8    `json:"category_id" validate:"required,numeric"`
	Pictures   []string `json:"pictures"`
}

type UpdateProductRequest struct {
	ID         uint64   `json:"id" validate:"required,numeric"`
	Name       string   `json:"name" validate:"max=150"`
	Desc       string   `json:"desc" validate:"max=1000"`
	SKU        string   `json:"sku"`
	Stock      uint16   `json:"stock" validate:"numeric,max=999"`
	Price      uint32   `json:"price" validate:"numeric,max=1000000000"`
	CategoryID uint8    `json:"category_id" validate:"numeric"`
	Pictures   []string `json:"pictures"`
}

type GetProductsRequest struct {
	CategoryID uint8 `json:"category_id,omitempty" validate:"numeric"`
	Page       uint8 `json:"page" validate:"required,numeric,min=1"`
	Size       uint8 `json:"size" validate:"required,numeric,min=1"`
}

type GetProductResponse struct {
	ID           uint64   `json:"id"`
	Name         string   `json:"name"`
	Desc         string   `json:"desc"`
	SKU          string   `json:"sku"`
	Stock        uint16   `json:"stock"`
	Price        uint32   `json:"price"`
	Pictures     []string `json:"pictures"`
	CategoryID   uint8    `json:"category_id"`
	CategoryName string   `json:"category_name"`
}
