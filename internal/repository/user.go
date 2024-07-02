package repository

import (
	"gorm.io/gorm"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/model"
)

type UserRepository struct {
	Repository[model.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repository: Repository[model.User]{
			DB: db,
		},
	}
}

func (r *UserRepository) CountByEmail(email string) (int64, error) {
	var total int64
	err := r.DB.Model(new(model.User)).Where("email = ?", email).Count(&total).Error
	return total, err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var result *model.User
	err := r.DB.Where("email = ?", email).First(&result).Error
	return result, err
}

func (r *UserRepository) FindByToken(token string) (*model.User, error) {
	var result *model.User
	err := r.DB.Where("token = ?", token).First(&result).Error
	return result, err
}
