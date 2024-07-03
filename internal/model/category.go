package model

type Category struct {
	ID       uint8     `gorm:"column:id;primaryKey"`
	Name     string    `gorm:"column:name"`
	Products []Product `gorm:"foreignKey:category_id"`
}

func (c *Category) TableName() string {
	return "category"
}
