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

func (r *UserRepository) FindByEmail(user *model.User, email string) error {
	return r.DB.Where("email = ?", email).First(user).Error
}

func (r *UserRepository) FindByToken(user *model.User, token string) error {
	return r.DB.Where("token = ?", token).First(user).Error
}
