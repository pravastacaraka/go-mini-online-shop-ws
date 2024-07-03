package model

type Cart struct {
	ID          uint64        `json:"id,omitempty" gorm:"column:id;primaryKey"`
	UserID      uint64        `json:"-" gorm:"column:user_id"`
	User        *User         `json:"user,omitempty" gorm:"foreignKey:user_id"`
	CartDetails []*CartDetail `json:"cart_details,omitempty" gorm:"foreignKey:cart_id"`
}

func (c *Cart) TableName() string {
	return "cart"
}
