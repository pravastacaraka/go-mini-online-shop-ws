package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiber(cfg *viper.Viper) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      cfg.GetString("app.name"),
		ErrorHandler: newErrorHandler(),
		Prefork:      cfg.GetBool("web.prefork"),
	})

	return app
}

func newErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		errorMsg := err.Error()

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		switch code {
		case fiber.StatusUnauthorized:
			if errorMsg == fiber.ErrUnauthorized.Message {
				errorMsg = "User is not authorized!"
			}
		case fiber.StatusBadRequest:
			if errorMsg == fiber.ErrBadRequest.Message {
				errorMsg = "Something wrong with the request, please check!"
			}
		case fiber.StatusInternalServerError:
			if errorMsg == fiber.ErrInternalServerError.Message {
				errorMsg = "There is something wrong, please try again!"
			}
		}

		return ctx.Status(code).JSON(fiber.Map{
			"error": errorMsg,
		})
	}
}
