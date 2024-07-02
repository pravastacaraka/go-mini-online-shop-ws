package model

import "time"

type User struct {
	ID        uint64    `gorm:"column:id;primaryKey"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	Token     string    `gorm:"column:token"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Addresses []Address `gorm:"foreignKey:user_id;references:id"`
}

func (u *User) TableName() string {
	return "user"
}
