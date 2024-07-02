package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type RouteConfig struct {
	App    *fiber.App
	DB     *gorm.DB
	Config *viper.Viper
}

func (c *RouteConfig) Setup() {
	c.App.Use(recover.New())

	c.SetupRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupRoute() {
	c.App.Post("/api/users", func(c *fiber.Ctx) error {
		return c.SendString("hello register users")
	})
	c.App.Get("/api/users/_login", func(c *fiber.Ctx) error {
		return c.SendString("hello login user")
	})
}

func (c *RouteConfig) SetupAuthRoute() {

}
