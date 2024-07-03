package controller

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/domain"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/usecase"
)

type ProductController struct {
	UseCase *usecase.ProductUseCase
}

func NewProductController(useCase *usecase.ProductUseCase) *ProductController {
	return &ProductController{
		UseCase: useCase,
	}
}

func (c *ProductController) Add(ctx *fiber.Ctx) error {
	request := new(domain.AddProductRequest)

	if err := ctx.BodyParser(request); err != nil {
		log.Errorf("failed to parse request body, err: %s", err.Error())
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Add(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(domain.WebResponse[*domain.CommonProductResponse]{Data: response})
}

func (c *ProductController) Get(ctx *fiber.Ctx) error {
	productID, _ := strconv.Atoi(ctx.Params("productId"))

	response, err := c.UseCase.Get(ctx.UserContext(), uint64(productID))
	if err != nil {
		return err
	}

	return ctx.JSON(domain.WebResponse[*domain.GetProductResponse]{Data: response})
}

func (c *ProductController) Update(ctx *fiber.Ctx) error {
	productID, _ := strconv.Atoi(ctx.Params("productId"))

	request := &domain.UpdateProductRequest{
		ID: uint64(productID),
	}

	if err := ctx.BodyParser(request); err != nil {
		log.Errorf("failed to parse request body, err: %s", err.Error())
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(domain.WebResponse[*domain.CommonProductResponse]{Data: response})
}

func (c *ProductController) List(ctx *fiber.Ctx) error {
	request := &domain.GetProductsRequest{
		CategoryID: uint8(ctx.QueryInt("category_id", 0)),
		Page:       uint8(ctx.QueryInt("page", 1)),
		Size:       uint8(ctx.QueryInt("size", 1)),
	}

	responses, total, err := c.UseCase.List(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	paging := &domain.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: uint16(total),
		TotalPage: uint16(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(domain.WebResponse[[]*domain.GetProductResponse]{
		Data:   responses,
		Paging: paging,
	})
}
