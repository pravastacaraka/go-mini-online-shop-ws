package usecase

import (
	"context"
	"strings"

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

// TODO: need changes with JWT implementation
func (c *UserUseCase) Verify(ctx context.Context, request *domain.AuthUserRequest) (*domain.AuthUserResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return nil, fiber.ErrBadRequest
	}

	token, err := c.UserRepo.FindByID(request.ID)
	if err != nil {
		log.Errorf("failed to get token by id, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if token != request.Token {
		return nil, fiber.ErrUnauthorized
	}

	return &domain.AuthUserResponse{ID: request.ID}, nil
}

func (c *UserUseCase) Create(ctx context.Context, request *domain.RegisterUserRequest) (*domain.RegisterUserResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return nil, fiber.ErrBadRequest
	}

	request.Email = strings.ToLower(request.Email)

	total, err := c.UserRepo.CountByEmail(request.Email)
	if err != nil {
		log.Errorf("failed to count user by email, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}
	if total > 0 {
		return nil, fiber.NewError(fiber.StatusConflict, "Email has been registered")
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
		Address:    strings.ToTitle(request.Address),
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

	request.Email = strings.ToLower(request.Email)

	user, err := c.UserRepo.FindByEmail(request.Email)
	if err != nil {
		log.Errorf("failed to get user by email, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}
	if user == nil || user.ID < 1 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "User is not registered")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		log.Errorf("failed to compare encrypted password, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
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
