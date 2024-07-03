package model

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"
)

type Product struct {
	ID         uint64      `gorm:"column:id;primaryKey"`
	Name       string      `gorm:"column:name"`
	Desc       string      `gorm:"column:desc"`
	SKU        string      `gorm:"column:sku"`
	Stock      uint16      `gorm:"column:stock"`
	Price      uint32      `gorm:"column:price"`
	Pictures   MultiString `gorm:"column:pictures;type:text"`
	CreatedAt  time.Time   `gorm:"column:created_at"`
	UpdatedAt  time.Time   `gorm:"column:updated_at"`
	CategoryID uint8       `gorm:"column:category_id" json:"-"`
	Category   Category    `gorm:"foreignKey:category_id"`
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
