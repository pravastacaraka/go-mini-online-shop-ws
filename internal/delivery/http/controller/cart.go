package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/delivery/http/middleware"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/domain"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/usecase"
)

type CartController struct {
	UseCase *usecase.CartUseCase
}

func NewCartController(useCase *usecase.CartUseCase) *CartController {
	return &CartController{
		UseCase: useCase,
	}
}

func (c *CartController) Add(ctx *fiber.Ctx) error {
	auth := middleware.GetAuthenticatedUserID(ctx)

	request := &domain.AddToCartRequest{
		UserID: auth.ID,
	}

	if err := ctx.BodyParser(request); err != nil {
		log.Errorf("failed to parse request body, err: %s", err.Error())
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Add(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(domain.WebResponse[*domain.CommonCartResponse]{Data: response})
}

func (c *CartController) Delete(ctx *fiber.Ctx) error {
	cartID, _ := strconv.Atoi(ctx.Params("cartID"))

	request := &domain.DeleteCartRequest{
		CartID: uint64(cartID),
	}

	response, err := c.UseCase.Delete(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(domain.WebResponse[*domain.CommonCartResponse]{Data: response})
}

func (c *CartController) UpdateDetail(ctx *fiber.Ctx) error {
	cartDetailID, _ := strconv.Atoi(ctx.Params("cartDetailID"))

	request := &domain.UpdateCartDetailRequest{
		CartDetailID: uint64(cartDetailID),
	}

	if err := ctx.BodyParser(request); err != nil {
		log.Errorf("failed to parse request body, err: %s", err.Error())
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.UpdateDetail(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(domain.WebResponse[*domain.CommonCartResponse]{Data: response})
}

func (c *CartController) DeleteDetail(ctx *fiber.Ctx) error {
	cartDetailID, _ := strconv.Atoi(ctx.Params("cartDetailID"))

	request := &domain.DeleteCartDetailRequest{
		CartDetailID: uint64(cartDetailID),
	}

	response, err := c.UseCase.DeleteDetail(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(domain.WebResponse[*domain.CommonCartResponse]{Data: response})
}

func (c *CartController) List(ctx *fiber.Ctx) error {
	auth := middleware.GetAuthenticatedUserID(ctx)

	request := &domain.GetCartListRequest{
		UserID: auth.ID,
	}

	response, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(domain.WebResponse[*domain.GetCartListResponse]{Data: response})
}

func (c *CartController) Checkout(ctx *fiber.Ctx) error {
	auth := middleware.GetAuthenticatedUserID(ctx)
	cartID, _ := strconv.Atoi(ctx.Params("cartID"))

	request := &domain.CheckoutRequest{
		UserID: auth.ID,
		CartID: uint64(cartID),
	}

	response, err := c.UseCase.Checkout(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(domain.WebResponse[*domain.CheckoutResponse]{Data: response})
}
