package model

import (
	"fmt"
	"time"
)

type User struct {
	ID        uint64     `json:"id,omitempty" gorm:"column:id;primaryKey"`
	Email     string     `json:"email,omitempty" gorm:"column:email"`
	Password  string     `json:"password,omitempty" gorm:"column:password"`
	Name      string     `json:"name,omitempty" gorm:"column:name"`
	Token     string     `json:"token,omitempty" gorm:"column:token"`
	CreatedAt time.Time  `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" gorm:"column:updated_at"`
	Addresses []*Address `json:"addresses,omitempty" gorm:"foreignKey:user_id"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) GetRedisKey(id any) string {
	return fmt.Sprintf("user:%d", id)
}

func (u *User) GetRedisTTL() time.Duration {
	return 24 * time.Hour
}
