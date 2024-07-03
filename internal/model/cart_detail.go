package model

type CartDetail struct {
	ID        uint64   `json:"id,omitempty" gorm:"column:id;primaryKey"`
	CartID    uint64   `json:"cart_id,omitempty" gorm:"column:cart_id"`
	Quantity  uint16   `json:"quantity,omitempty" gorm:"column:quantity"`
	ProductID uint64   `json:"-" gorm:"column:product_id"`
	Product   *Product `json:"product,omitempty" gorm:"foreignKey:product_id"`
}

func (c *CartDetail) TableName() string {
	return "cart_detail"
}
