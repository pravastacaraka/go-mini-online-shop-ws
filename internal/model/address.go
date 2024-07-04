package model

import "time"

type Address struct {
	ID         uint64    `json:"id,omitempty" gorm:"column:id;primaryKey"`
	UserID     uint64    `json:"user_id,omitempty" gorm:"column:user_id"`
	Address    string    `json:"address_id,omitempty" gorm:"column:address"`
	PostalCode uint32    `json:"postal_code,omitempty" gorm:"column:postal_code"`
	CreatedAt  time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (u *Address) TableName() string {
	return "address"
}
