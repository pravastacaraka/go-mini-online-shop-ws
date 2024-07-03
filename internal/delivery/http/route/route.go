package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/delivery/http/controller"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *controller.UserController
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
	v1 := c.App.Group("/api/v1/users")
	v1.Post("/login", c.UserController.Login)
	v1.Post("/register", c.UserController.Register)
}

func (c *RouteConfig) SetupAuthRoutes() {

}
