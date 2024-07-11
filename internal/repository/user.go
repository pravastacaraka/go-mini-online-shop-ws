package repository

import (
	"encoding/json"

	"gorm.io/gorm"

	"github.com/gofiber/storage/redis/v3"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/model"
)

type UserRepository struct {
	Repository[model.User]
}

func NewUserRepository(db *gorm.DB, redis *redis.Storage) *UserRepository {
	return &UserRepository{
		Repository: Repository[model.User]{
			DB:    db,
			Redis: redis,
		},
	}
}

func (r *UserRepository) Create(tx *gorm.DB, data *model.User) error {
	db := r.DB
	if tx != nil {
		db = tx
	}

	if err := db.Create(data).Error; err != nil {
		return err
	}

	jsn, _ := json.Marshal(data)
	if err := r.Redis.Set(data.GetRedisKey(data.ID), jsn, data.GetRedisTTL()); err != nil {
		return nil
	}

	return nil
}

func (r *UserRepository) Update(tx *gorm.DB, data *model.User) error {
	db := r.DB
	if tx != nil {
		db = tx
	}

	if err := db.Save(data).Error; err != nil {
		return err
	}

	jsn, _ := json.Marshal(data)
	if err := r.Redis.Set(data.GetRedisKey(data.ID), jsn, data.GetRedisTTL()); err != nil {
		return nil
	}

	return nil
}

func (r *UserRepository) FindByID(id any) (string, error) {
	var (
		err  error
		user *model.User
	)

	result, err := r.Redis.Get(user.GetRedisKey(id))
	if err != nil || len(result) < 1 {
		if err = r.DB.First(&user, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return "", nil
			} else {
				return "", err
			}
		}

		jsn, _ := json.Marshal(user)
		if err = r.Redis.Set(user.GetRedisKey(id), jsn, user.GetRedisTTL()); err != nil {
			return "", nil
		}
	}

	if err = json.Unmarshal(result, &user); err != nil {
		return "", nil
	}

	if user == nil || user.ID < 1 {
		return "", nil
	}

	return user.Token, err
}

func (r *UserRepository) CountByEmail(email string) (int64, error) {
	var total int64
	err := r.DB.Model(&model.User{}).Where("email = ?", email).Count(&total).Error
	return total, err
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var (
		err  error
		user *model.User
	)

	if err = r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, nil
		} else {
			return user, err
		}
	}

	jsn, _ := json.Marshal(user)
	if err = r.Redis.Set(user.GetRedisKey(user.ID), jsn, user.GetRedisTTL()); err != nil {
		return user, nil
	}

	return user, nil
}
