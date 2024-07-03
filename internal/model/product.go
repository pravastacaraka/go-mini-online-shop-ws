package model

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"
)

type Product struct {
	ID         uint64      `json:"id,omitempty" gorm:"column:id;primaryKey"`
	Name       string      `json:"name,omitempty" gorm:"column:name"`
	Desc       string      `json:"desc,omitempty" gorm:"column:desc"`
	SKU        string      `json:"sku,omitempty" gorm:"column:sku"`
	Stock      uint16      `json:"stock,omitempty" gorm:"column:stock"`
	Price      uint32      `json:"price,omitempty" gorm:"column:price"`
	Pictures   MultiString `json:"pictures,omitempty" gorm:"column:pictures;type:text"`
	CreatedAt  time.Time   `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt  time.Time   `json:"updated_at,omitempty" gorm:"column:updated_at"`
	CategoryID uint8       `json:"-" gorm:"column:category_id"`
	Category   *Category   `json:"category,omitempty" gorm:"foreignKey:category_id"`
}

func (p *Product) TableName() string {
	return "product"
}

type MultiString []string

func (m MultiString) Value() (driver.Value, error) {
	return "{" + strings.Join(m, ",") + "}", nil
}

func (m *MultiString) Scan(value interface{}) error {
	if value == nil {
		*m = MultiString{}
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return errors.New("failed to scan []string")
	}

	str = strings.Trim(str, "{}")
	if str == "" {
		*m = MultiString{}
		return nil
	}

	*m = strings.Split(str, ",")
	return nil
}
