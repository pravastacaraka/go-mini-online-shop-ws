package repository

import (
	"github.com/gofiber/storage/redis/v3"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB    *gorm.DB
	Redis *redis.Storage
}

func (r *Repository[T]) Create(tx *gorm.DB, data *T) error {
	db := r.DB
	if tx != nil {
		db = tx
	}
	return db.Create(data).Error
}

func (r *Repository[T]) Update(tx *gorm.DB, data *T) error {
	db := r.DB
	if tx != nil {
		db = tx
	}
	return db.Save(data).Error
}

func (r *Repository[T]) Delete(tx *gorm.DB, data *T) error {
	db := r.DB
	if tx != nil {
		db = tx
	}
	return db.Delete(data).Error
}

func (r *Repository[T]) FindByID(id any) (*T, error) {
	var data *T
	err := r.DB.First(&data, id).Error
	return data, err
}
