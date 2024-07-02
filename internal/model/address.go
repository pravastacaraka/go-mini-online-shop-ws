package model

import "time"

type Address struct {
	ID         uint64    `gorm:"column:id;primaryKey"`
	UserID     uint64    `gorm:"column:user_id"`
	Address    string    `gorm:"column:address"`
	PostalCode uint32    `gorm:"column:postal_code"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (u *Address) TableName() string {
	return "address"
}
