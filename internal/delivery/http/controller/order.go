package controller

import (
	"math"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/delivery/http/middleware"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/domain"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/usecase"
)

type OrderController struct {
	UseCase *usecase.OrderUseCase
}

func NewOrderController(useCase *usecase.OrderUseCase) *OrderController {
	return &OrderController{
		UseCase: useCase,
	}
}

func (c *OrderController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetAuthenticatedUserID(ctx)

	request := &domain.CreateOrderRequest{
		Customer: domain.Customer{
			UserID: auth.ID,
		},
	}

	if err := ctx.BodyParser(request); err != nil {
		log.Errorf("failed to parse request body, err: %s", err.Error())
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(domain.WebResponse[*domain.CreateOrderResponse]{Data: response})
}

func (c *OrderController) List(ctx *fiber.Ctx) error {
	auth := middleware.GetAuthenticatedUserID(ctx)

	now := time.Now()

	startDate, err := time.Parse("2006-01-02", ctx.Params("start_date", now.AddDate(0, -3, 0).Format("2006-01-02")))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "start date format is wrong, it should be 2006-01-02")
	}

	endDate, err := time.Parse("2006-01-02", ctx.Params("end_date", now.AddDate(0, 0, 1).Format("2006-01-02")))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "end date format is wrong, it should be 2006-01-02")
	}

	request := &domain.GetOrdersRequest{
		UserID:    auth.ID,
		StartDate: startDate,
		EndDate:   endDate,
		Invoice:   ctx.Query("invoice"),
		Page:      uint8(ctx.QueryInt("page", 1)),
		Size:      uint8(ctx.QueryInt("size", 10)),
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

	return ctx.JSON(domain.WebResponse[[]*domain.GetOrderBuyerResponse]{
		Data:   responses,
		Paging: paging,
	})
}

func (c *OrderController) Pay(ctx *fiber.Ctx) error {
	auth := middleware.GetAuthenticatedUserID(ctx)

	paymentID, _ := strconv.Atoi(ctx.Params("paymentId"))

	request := &domain.DoPaymentOrderRequest{
		UserID:    auth.ID,
		PaymentID: uint64(paymentID),
	}

	if err := ctx.BodyParser(request); err != nil {
		log.Errorf("failed to parse request body, err: %s", err.Error())
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.DoPayment(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(domain.WebResponse[*domain.DoPaymentOrderResponse]{Data: response})
}
