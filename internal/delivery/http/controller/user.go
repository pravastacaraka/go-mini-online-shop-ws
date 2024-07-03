package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/domain"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/usecase"
)

type UserController struct {
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase) *UserController {
	return &UserController{
		UseCase: useCase,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(domain.RegisterUserRequest)

	if err := ctx.BodyParser(request); err != nil {
		log.Errorf("failed to parse request body, err: %s", err.Error())
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		log.Errorf("failed to register user, err: %s", err.Error())
		return err
	}

	return ctx.JSON(domain.WebResponse[*domain.RegisterUserResponse]{Data: response})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(domain.LoginUserRequest)

	if err := ctx.BodyParser(request); err != nil {
		log.Errorf("failed to parse request body, err: %s", err.Error())
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Login(ctx.UserContext(), request)
	if err != nil {
		log.Errorf("failed to login user, err: %s", err.Error())
		return err
	}

	return ctx.JSON(domain.WebResponse[*domain.LoginUserResponse]{Data: response})
}
