package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/delivery/http/controller"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/delivery/http/middleware"
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
	productRepo := repository.NewProductRepository(config.DB)
	cartRepo := repository.NewCartRepository(config.DB)
	cartDetailRepo := repository.NewCartDetailRepository(config.DB)

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Validator, userRepo, addressRepo)
	productUseCase := usecase.NewProductUseCase(config.DB, config.Validator, productRepo)
	cartUseCase := usecase.NewCartUseCase(config.DB, config.Validator, productRepo, cartRepo, cartDetailRepo)

	// setup controllers
	userController := controller.NewUserController(userUseCase)
	productController := controller.NewProductController(productUseCase)
	cartController := controller.NewCartController(cartUseCase)

	// setup custom middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	route := route.RouteConfig{
		App:               config.App,
		AuthMiddleware:    authMiddleware,
		UserController:    userController,
		ProductController: productController,
		CartController:    cartController,
	}
	route.Setup()
}
