package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/delivery/http/controller"
)

type RouteConfig struct {
	App               *fiber.App
	AuthMiddleware    fiber.Handler
	UserController    *controller.UserController
	ProductController *controller.ProductController
}

func (c *RouteConfig) Setup() {
	c.SetupMiddlewares()
	c.SetupRoutes()
	c.SetupAuthRoutes()
}

func (c *RouteConfig) SetupMiddlewares() {
	loggerCfg := logger.Config{
		TimeFormat: "02 Jan 2006 15:04:05",
	}

	// Add default middlewares
	c.App.Use(
		recover.New(),
		logger.New(loggerCfg),
	)
}

func (c *RouteConfig) SetupRoutes() {
	v1 := c.App.Group("/api/v1")

	users := v1.Group("/users")
	users.Post("/login", c.UserController.Login)
	users.Post("/register", c.UserController.Register)

	products := v1.Group("/products")
	products.Get("/", c.ProductController.List)
	products.Get("/:productId", c.ProductController.Get)
}

func (c *RouteConfig) SetupAuthRoutes() {
	c.App.Use(c.AuthMiddleware)

	v1 := c.App.Group("/api/v1")

	products := v1.Group("/products")
	products.Post("/", c.ProductController.Add)
	products.Patch("/:productId", c.ProductController.Update)

	cart := v1.Group("/cart")
	cart.Get("/list", func(c *fiber.Ctx) error {
		return c.SendString("ini isi cart kamu")
	})
}
