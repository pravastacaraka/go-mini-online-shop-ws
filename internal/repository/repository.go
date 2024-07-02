package repository

import "gorm.io/gorm"

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(data *T) error {
	return r.DB.Create(data).Error
}

func (r *Repository[T]) Update(data *T) error {
	return r.DB.Save(data).Error
}

func (r *Repository[T]) Delete(data *T) error {
	return r.DB.Delete(data).Error
}

func (r *Repository[T]) FindByID(data *T, id any) error {
	return r.DB.Where("id = ?", id).Take(data).Error
}
