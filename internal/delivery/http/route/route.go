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
	CartController    *controller.CartController
	OrderController   *controller.OrderController
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
	products.Post("/add", c.ProductController.Add)
	products.Patch("/:productId", c.ProductController.Update)

	cart := v1.Group("/cart")
	cart.Get("/", c.CartController.List)
	cart.Post("/add", c.CartController.Add)
	cart.Delete("/:cartId", c.CartController.Delete)
	cart.Patch("/:cartDetailId", c.CartController.UpdateDetail)
	cart.Delete("/:cartDetailId", c.CartController.DeleteDetail)

	checkout := v1.Group("/checkout")
	checkout.Get("/:cartId", c.CartController.Checkout)

	order := v1.Group("/order")
	order.Get("/", c.OrderController.List)
	order.Post("/create", c.OrderController.Create)
	order.Get("/:orderId", c.OrderController.Get)
	order.Post("/pay/:paymentId", c.OrderController.Pay)
}
