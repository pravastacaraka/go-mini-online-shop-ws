package repository

import (
	"gorm.io/gorm"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/model"
)

type AddressRepository struct {
	Repository[model.Address]
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{
		Repository: Repository[model.Address]{
			DB: db,
		},
	}
}
