package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/delivery/http/controller"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/delivery/http/route"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/repository"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/usecase"
)

type BootstrapConfig struct {
	Config    *viper.Viper
	DB        *gorm.DB
	Validator *validator.Validate
	App       *fiber.App
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepo := repository.NewUserRepository(config.DB)
	addressRepo := repository.NewAddressRepository(config.DB)

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.Validator, userRepo, addressRepo)

	// setup controllers
	userController := controller.NewUserController(userUseCase)

	route := route.RouteConfig{
		App:            config.App,
		UserController: userController,
	}
	route.Setup()
}
