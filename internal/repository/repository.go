package repository

import "gorm.io/gorm"

type Repository[T any] struct {
	DB *gorm.DB
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

func (r *Repository[T]) FindByID(data *T, id any) error {
	return r.DB.Where("id = ?", id).Take(data).Error
}
