package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/domain"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/model"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/repository"
)

type UserUseCase struct {
	DB          *gorm.DB
	Validate    *validator.Validate
	UserRepo    *repository.UserRepository
	AddressRepo *repository.AddressRepository
}

func NewUserUseCase(
	db *gorm.DB,
	validate *validator.Validate,
	userRepo *repository.UserRepository,
	addressRepo *repository.AddressRepository,
) *UserUseCase {
	return &UserUseCase{
		DB:          db,
		Validate:    validate,
		UserRepo:    userRepo,
		AddressRepo: addressRepo,
	}
}

func (c *UserUseCase) Create(ctx context.Context, request *domain.RegisterUserRequest) (*domain.RegisterUserResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return nil, fiber.ErrBadRequest
	}

	total, err := c.UserRepo.CountByEmail(request.Email)
	if err != nil {
		log.Errorf("failed to count user by email, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}
	if total > 0 {
		log.Error("user already exists!")
		return nil, fiber.ErrConflict
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("failed to encrypt password, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	tx := c.DB.WithContext(ctx).Begin()
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
		log.Errorf("failed to commit tx, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &domain.RegisterUserResponse{ID: user.ID}, nil
}

func (c *UserUseCase) Login(ctx context.Context, request *domain.LoginUserRequest) (*domain.LoginUserResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return nil, fiber.ErrBadRequest
	}

	user, err := c.UserRepo.FindByEmail(request.Email)
	if err != nil {
		log.Errorf("failed to get user by email, err: %s", err.Error())
		return nil, fiber.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		log.Errorf("failed to compare encrypted password, err: %s", err.Error())
		return nil, fiber.ErrUnauthorized
	}

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// TODO: Need changes with JWT implementation
	user.Token = uuid.New().String()
	if err := c.UserRepo.Update(tx, user); err != nil {
		log.Errorf("failed to update user session, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Errorf("failed to commit tx, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &domain.LoginUserResponse{ID: user.ID, Token: user.Token}, nil
}
