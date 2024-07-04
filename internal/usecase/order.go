package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/domain"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/model"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/repository"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/utils"
)

type OrderUseCase struct {
	DB              *gorm.DB
	Validate        *validator.Validate
	ProductRepo     *repository.ProductRepository
	CartRepo        *repository.CartRepository
	PaymentRepo     *repository.PaymentRepository
	OrderRepo       *repository.OrderRepository
	OrderDetailRepo *repository.OrderDetailRepository
}

func NewOrderUseCase(
	db *gorm.DB,
	validate *validator.Validate,
	productRepo *repository.ProductRepository,
	cartRepo *repository.CartRepository,
	paymentRepo *repository.PaymentRepository,
	orderRepo *repository.OrderRepository,
	orderDetailRepo *repository.OrderDetailRepository,
) *OrderUseCase {
	return &OrderUseCase{
		DB:              db,
		Validate:        validate,
		CartRepo:        cartRepo,
		ProductRepo:     productRepo,
		PaymentRepo:     paymentRepo,
		OrderRepo:       orderRepo,
		OrderDetailRepo: orderDetailRepo,
	}
}

func (c *OrderUseCase) Create(ctx context.Context, request *domain.CreateOrderRequest) (*domain.CreateOrderResponse, error) {
	var result *domain.CreateOrderResponse

	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return result, fiber.ErrBadRequest
	}

	// Check cart, is it available? If yes then proceed the order, otherwise declined
	countCart, err := c.CartRepo.CountByID(request.CartID)
	if err != nil {
		log.Errorf("failed to count cart, err: %s", err.Error())
		return result, fiber.ErrInternalServerError
	}
	if countCart < 1 {
		return result, fiber.NewError(fiber.StatusBadRequest, "You don't have any cart. Please add your item(s) to your cart!")
	}

	// Check payment, does the order already created? If no then proceed the order, otherwise declined
	countPayment, err := c.PaymentRepo.CountByCartID(request.CartID)
	if err != nil {
		log.Errorf("failed to count payment, err: %s", err.Error())
		return result, fiber.ErrInternalServerError
	}
	if countPayment > 0 {
		return result, fiber.NewError(fiber.StatusBadRequest, "Your order has been created previously. Please make a payment!")
	}

	// Check ordered quantity with last product stock
	var (
		productIDs  []uint64
		mapProducts = make(map[uint64]*domain.OrderDetail)
	)
	for _, val := range request.OrderDetails {
		mapProducts[val.ProductID] = val
		productIDs = append(productIDs, val.ProductID)
	}
	products, err := c.ProductRepo.FindByIDs(productIDs)
	if err != nil {
		log.Errorf("failed to get products, err: %s", err.Error())
		return result, fiber.ErrInternalServerError
	}
	for _, product := range products {
		val, ok := mapProducts[product.ID]
		if !ok {
			continue
		}
		if val.Quantity <= product.Stock {
			continue
		}
		log.Errorf(`Product "%s" is out of stock`, val.ProductName)
		return result, fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf(`Product "%s" is out of stock`, val.ProductName))
	}

	invoice := utils.GenerateInvoice()

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// Insert payment
	payment := &model.Payment{
		CartID:      request.CartID,
		Amount:      request.Payment.TotalPayment,
		GatewayName: "Saldo", // Still harcoded
		Status:      model.PaymentStatusPending,
		PaymentCode: uuid.New().String(),
	}
	if err := c.PaymentRepo.Create(tx, payment); err != nil {
		log.Errorf("failed to create payment, err: %s", err.Error())
		return result, fiber.ErrInternalServerError
	}

	// Insert order
	order := &model.Order{
		PaymentID:     payment.ID,
		UserID:        request.Customer.UserID,
		AddressID:     request.Customer.AddressID,
		Invoice:       invoice,
		TotalPrice:    request.Payment.TotalPrice,
		TotalWeight:   request.Shipment.ShippingWeight,
		ShippingPrice: request.Shipment.ShippingPrice,
		Status:        model.OrderStatusCreated,
	}
	if err := c.OrderRepo.Create(tx, order); err != nil {
		log.Errorf("failed to create order, err: %s", err.Error())
		return result, fiber.ErrInternalServerError
	}

	// Insert order details
	orderDetails := []*model.OrderDetail{}
	for _, val := range request.OrderDetails {
		temp := &model.OrderDetail{
			OrderID:        order.ID,
			ProductID:      val.ProductID,
			ProductName:    val.ProductName,
			Quantity:       val.Quantity,
			Price:          val.Price,
			SubTotalPrice:  val.SubTotalPrice,
			Weight:         val.Weight,
			SubTotalWeight: val.SubTotalWeight,
			CategoryID:     val.CategoryID,
		}
		orderDetails = append(orderDetails, temp)
	}
	if err := c.OrderDetailRepo.BulkCreate(tx, orderDetails); err != nil {
		log.Errorf("failed to create order, err: %s", err.Error())
		return result, fiber.ErrInternalServerError
	}

	// Delete cart after order created, because we don't need anymore
	if err := c.CartRepo.DeleteByID(tx, payment.CartID); err != nil {
		log.Errorf("failed to delete cart, err: %s", err.Error())
		return result, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Errorf("failed to commit tx, err: %s", err.Error())
		return result, fiber.ErrInternalServerError
	}

	result = &domain.CreateOrderResponse{
		Invoice:     invoice,
		PaymentID:   payment.ID,
		PaymentCode: payment.PaymentCode,
		Message:     "Order has been created successfully. Please make a payment first!",
	}

	return result, nil
}

func (c *OrderUseCase) DoPayment(ctx context.Context, request *domain.DoPaymentOrderRequest) (*domain.DoPaymentOrderResponse, error) {
	var result *domain.DoPaymentOrderResponse

	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return result, fiber.ErrBadRequest
	}

	// Check payment, does the pending payment is exist? If yes proceed payment, otherwise declined
	payment, err := c.PaymentRepo.FindByID(request.PaymentID)
	if err != nil {
		if !strings.Contains(err.Error(), "no record") {
			log.Errorf("failed to get payment, err: %s", err.Error())
			return result, fiber.ErrInternalServerError
		}
	}
	if payment == nil || payment.ID < 1 {
		return result, fiber.NewError(fiber.StatusBadRequest, "Payment does not found")
	}
	if request.PaymentCode != payment.PaymentCode {
		return result, fiber.NewError(fiber.StatusBadRequest, "Payment code does not same")
	}

	// Check order, does the pending order is exist? If yes proceed payment, otherwise declined
	order, err := c.OrderRepo.FindByPaymentID(request.PaymentID)
	if err != nil {
		if !strings.Contains(err.Error(), "no record") {
			log.Errorf("failed to get order, err: %s", err.Error())
			return result, fiber.ErrInternalServerError
		}
	}
	if order == nil || order.ID < 1 {
		return result, fiber.NewError(fiber.StatusBadRequest, "Order with this payment does not found")
	}

	// Check ordered quantity with last product stock
	var (
		productIDs      []uint64
		orderedProducts = make(map[uint64]*model.OrderDetail)
	)
	for _, val := range order.OrderDetails {
		orderedProducts[val.ProductID] = val
		productIDs = append(productIDs, val.ProductID)
	}
	products, err := c.ProductRepo.FindByIDs(productIDs)
	if err != nil {
		log.Errorf("failed to get products, err: %s", err.Error())
		return result, fiber.ErrInternalServerError
	}
	for _, product := range products {
		val, ok := orderedProducts[product.ID]
		if !ok {
			continue
		}
		if val.Quantity <= product.Stock {
			continue
		}
		// TODO: Need to send events to cancel pending payment and order asynchronously
		log.Errorf(`Product "%s" is out of stock`, val.ProductName)
		return result, fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf(`Product "%s" is out of stock`, val.ProductName))
	}

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	payment.Status = model.PaymentStatusSucceed
	if err := c.PaymentRepo.Update(tx, payment); err != nil {
		log.Errorf("failed to update payment status, err: %s", err.Error())
		return result, fiber.ErrInternalServerError
	}

	order.Status = model.OrderStatusProcessed
	if err := c.OrderRepo.Update(tx, order); err != nil {
		log.Errorf("failed to update order status, err: %s", err.Error())
		return result, fiber.ErrInternalServerError
	}

	for _, product := range products {
		qty := orderedProducts[product.ID].Quantity
		tx.Model(&product).UpdateColumn("stock", gorm.Expr("stock - ?", qty))
	}

	if err := tx.Commit().Error; err != nil {
		log.Errorf("failed to commit tx, err: %s", err.Error())
		return result, fiber.ErrInternalServerError
	}

	result = &domain.DoPaymentOrderResponse{
		Invoice: order.Invoice,
		Message: "Payment has been successfully made",
	}

	return result, nil
}

func (c *OrderUseCase) Get(ctx context.Context, request *domain.GetOrderBuyerRequest) (*domain.GetOrderBuyerResponse, error) {
	var result *domain.GetOrderBuyerResponse

	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return result, fiber.ErrBadRequest
	}

	// TODO: Add get order detail use case

	return result, nil
}

func (c *OrderUseCase) List(ctx context.Context, request *domain.GetOrdersRequest) ([]*domain.GetOrderBuyerResponse, uint16, error) {
	var result []*domain.GetOrderBuyerResponse

	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return result, 0, fiber.ErrBadRequest
	}

	orders, total, err := c.OrderRepo.Find(request)
	if err != nil {
		log.Errorf("failed to get all orders, err: %s", err.Error())
		return result, 0, fiber.ErrInternalServerError
	}

	for _, order := range orders {
		var orderDetails []*domain.GetOrderBuyerResponse_OrderDetail

		for _, detail := range order.OrderDetails {
			temp := &domain.GetOrderBuyerResponse_OrderDetail{
				ProductName: detail.ProductName,
				Quantity:    detail.Quantity,
				Price:       utils.IDR(detail.Price),
				Weight:      detail.Weight,
				CategoryID:  detail.CategoryID,
			}
			orderDetails = append(orderDetails, temp)
		}

		temp := &domain.GetOrderBuyerResponse{
			ID:           order.ID,
			Invoice:      order.Invoice,
			Status:       model.OrderStatusDesc[order.Status],
			OrderDetails: orderDetails,
			TotalPrice:   utils.IDR(order.TotalPrice),
			Customer: domain.GetOrderBuyerResponse_Customer{
				Name:       order.User.Name,
				Address:    order.Address.Address,
				PostalCode: order.Address.PostalCode,
			},
			Shipment: domain.GetOrderBuyerResponse_Shipment{
				ShippingAgentName:    "Kurir Toko", // Still hardcoded
				ShippingAgentProduct: "Regular",
				ShippingPrice:        utils.IDR(order.ShippingPrice),
				ShippingWeight:       order.TotalWeight,
			},
			Payment: domain.GetOrderBuyerResponse_Payment{
				GatewayName:  order.Payment.GatewayName,
				TotalPayment: utils.IDR(order.Payment.Amount),
			},
		}
		result = append(result, temp)
	}

	return result, total, nil
}
