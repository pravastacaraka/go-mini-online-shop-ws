package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/domain"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/model"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/repository"
)

type UserUseCase struct {
	Validate    *validator.Validate
	UserRepo    *repository.UserRepository
	AddressRepo *repository.AddressRepository
}

func NewUserUseCase(
	validate *validator.Validate,
	userRepo *repository.UserRepository,
	addressRepo *repository.AddressRepository,
) *UserUseCase {
	return &UserUseCase{
		Validate:    validate,
		UserRepo:    userRepo,
		AddressRepo: addressRepo,
	}
}

func (c *UserUseCase) Create(ctx context.Context, request *domain.RegisterUserRequest) (*domain.UserResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return nil, fiber.ErrBadRequest
	}

	total, err := c.UserRepo.CountByEmail(request.Email)
	if err != nil {
		log.Errorf("failed to get count user, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}
	if total > 0 {
		log.Error("user already exists!")
		return nil, fiber.ErrConflict
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("failed to hash password, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	tx := c.UserRepo.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := &model.User{
		Email:    request.Email,
		Password: string(password),
		Name:     request.Name,
	}
	if err := c.UserRepo.Create(tx, user); err != nil {
		log.Errorf("failed to create user, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	address := &model.Address{
		UserID:     user.ID,
		Address:    request.Address,
		PostalCode: request.PostalCode,
	}
	if err := c.AddressRepo.Create(tx, address); err != nil {
		log.Errorf("failed to create address, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Errorf("failed commit tx, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &domain.UserResponse{ID: user.ID}, nil
}
