package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/delivery/http/controller"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *controller.UserController
}

func (c *RouteConfig) Setup() {
	c.App.Use(recover.New())

	c.SetupRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupRoute() {
	v1 := c.App.Group("/api/v1/users")
	v1.Get("/login", c.UserController.Login)
	v1.Post("/register", c.UserController.Register)
}

func (c *RouteConfig) SetupAuthRoute() {

}
