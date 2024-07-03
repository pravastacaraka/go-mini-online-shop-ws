package repository

import (
	"gorm.io/gorm"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/model"
)

type CartDetailRepository struct {
	Repository[model.CartDetail]
}

func NewCartDetailRepository(db *gorm.DB) *CartDetailRepository {
	return &CartDetailRepository{
		Repository: Repository[model.CartDetail]{
			DB: db,
		},
	}
}
