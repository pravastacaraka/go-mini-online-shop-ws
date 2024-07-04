package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"

	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/domain"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/model"
	"github.com/pravastacaraka/go-ws-mini-online-shop/internal/repository"
)

type CartUseCase struct {
	DB             *gorm.DB
	Validate       *validator.Validate
	ProductRepo    *repository.ProductRepository
	CartRepo       *repository.CartRepository
	CartDetailRepo *repository.CartDetailRepository
}

func NewCartUseCase(
	db *gorm.DB,
	validate *validator.Validate,
	productRepo *repository.ProductRepository,
	cartRepo *repository.CartRepository,
	cartDetailRepo *repository.CartDetailRepository,
) *CartUseCase {
	return &CartUseCase{
		DB:             db,
		Validate:       validate,
		ProductRepo:    productRepo,
		CartRepo:       cartRepo,
		CartDetailRepo: cartDetailRepo,
	}
}

func (c *CartUseCase) Add(ctx context.Context, request *domain.AddToCartRequest) (*domain.CommonCartResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return nil, fiber.ErrBadRequest
	}

	var message = "Your item has been added to your cart successfully!"

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	total, _ := c.CartRepo.CountByUserID(request.UserID)
	if total > 0 { // Update to existing cart detail
		product, err := c.ProductRepo.FindByID(request.ProductID)
		if err != nil {
			log.Errorf("failed to get product, err: %s", err.Error())
			return nil, fiber.ErrInternalServerError
		}
		if request.Quantity > product.Stock {
			log.Error("quantity must be lower than product stock!")
			return nil, fiber.ErrBadRequest
		}

		cart, err := c.CartRepo.FindByUserID(request.UserID)
		if err != nil {
			log.Infof("failed to get cart, err: %s", err.Error())
			return nil, fiber.ErrInternalServerError
		}

		mapProducts := make(map[uint64]*model.CartDetail)
		for _, val := range cart.CartDetails {
			if _, ok := mapProducts[val.ProductID]; !ok {
				mapProducts[val.ProductID] = val
			}
		}

		if cartDetail, ok := mapProducts[request.ProductID]; ok {
			if request.Quantity == cartDetail.Quantity {
				return nil, nil
			}
			cartDetail.Quantity = request.Quantity
			if err := c.CartDetailRepo.Update(tx, cartDetail); err != nil {
				log.Errorf("failed to update cart detail, err: %s", err.Error())
				return nil, fiber.ErrInternalServerError
			}
			message = "Your item has been updated to your cart successfully!"
		} else {
			cartDetail := &model.CartDetail{
				CartID:    cart.ID,
				ProductID: request.ProductID,
				Quantity:  request.Quantity,
			}
			if err := c.CartDetailRepo.Create(tx, cartDetail); err != nil {
				log.Errorf("failed to add cart detail, err: %s", err.Error())
				return nil, fiber.ErrInternalServerError
			}
		}
	} else { // Create new cart and cart detail
		product, err := c.ProductRepo.FindByID(request.ProductID)
		if err != nil {
			log.Errorf("failed to get product, err: %s", err.Error())
			return nil, fiber.ErrInternalServerError
		}
		if request.Quantity > product.Stock {
			log.Error("quantity must be lower than product stock!")
			return nil, fiber.ErrBadRequest
		}

		cart := &model.Cart{
			UserID: request.UserID,
		}
		if err := c.CartRepo.Create(tx, cart); err != nil {
			log.Errorf("failed to add cart, err: %s", err.Error())
			return nil, fiber.ErrInternalServerError
		}

		cartDetail := &model.CartDetail{
			CartID:    cart.ID,
			ProductID: request.ProductID,
			Quantity:  request.Quantity,
		}
		if err := c.CartDetailRepo.Create(tx, cartDetail); err != nil {
			log.Errorf("failed to add cart detail, err: %s", err.Error())
			return nil, fiber.ErrInternalServerError
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Errorf("failed to commit tx, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &domain.CommonCartResponse{Message: message}, nil
}

func (c *CartUseCase) Delete(ctx context.Context, request *domain.DeleteCartRequest) (*domain.CommonCartResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return nil, fiber.ErrBadRequest
	}

	cart, err := c.CartRepo.FindByID(request.CartID)
	if err != nil {
		log.Errorf("failed to get cart, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}
	if cart == nil {
		log.Errorf("cart id %d doesn't exist!", request.CartID)
		return nil, fiber.ErrNotFound
	}

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.CartRepo.Delete(tx, cart); err != nil {
		log.Errorf("failed to delete cart, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Errorf("failed to commit tx, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &domain.CommonCartResponse{Message: "Your cart has been deleted successfully!"}, nil
}

func (c *CartUseCase) UpdateDetail(ctx context.Context, request *domain.UpdateCartDetailRequest) (*domain.CommonCartResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return nil, fiber.ErrBadRequest
	}

	cartDetail, err := c.CartDetailRepo.FindByID(request.CartDetailID)
	if err != nil {
		log.Errorf("failed to get cart detail, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}
	if cartDetail == nil {
		log.Errorf("cart detail id %d doesn't exist!", request.CartDetailID)
		return nil, fiber.ErrNotFound
	}
	if request.Quantity == cartDetail.Quantity {
		return nil, nil
	}

	product, err := c.ProductRepo.FindByID(cartDetail.ProductID)
	if err != nil {
		log.Errorf("failed to get product, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}
	if request.Quantity > product.Stock {
		log.Error("quantity is higher than product stock!")
		return nil, fiber.ErrBadRequest
	}

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	cartDetail.Quantity = request.Quantity
	if err := c.CartDetailRepo.Update(tx, cartDetail); err != nil {
		log.Errorf("failed to update cart detail, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Errorf("failed to commit tx, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &domain.CommonCartResponse{Message: "Your item has been updated successfully!"}, nil
}

func (c *CartUseCase) DeleteDetail(ctx context.Context, request *domain.DeleteCartDetailRequest) (*domain.CommonCartResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return nil, fiber.ErrBadRequest
	}

	cartDetail, err := c.CartDetailRepo.FindByID(request.CartDetailID)
	if err != nil {
		log.Errorf("failed to get cart detail, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}
	if cartDetail == nil {
		log.Errorf("cart detail id %d doesn't exist!", request.CartDetailID)
		return nil, fiber.ErrNotFound
	}

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.CartDetailRepo.Delete(tx, cartDetail); err != nil {
		log.Errorf("failed to delete cart detail, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Errorf("failed to commit tx, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &domain.CommonCartResponse{Message: "Your item has been deleted successfully!"}, nil
}

func (c *CartUseCase) List(ctx context.Context, request *domain.GetCartListRequest) (*domain.GetCartListResponse, error) {
	var result *domain.GetCartListResponse

	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return result, fiber.ErrBadRequest
	}

	total, _ := c.CartRepo.CountByUserID(request.UserID)
	if total < 1 {
		return result, nil
	}

	cart, err := c.CartRepo.FindByUserID(request.UserID)
	if err != nil {
		log.Errorf("failed to get cart, err: %s", err.Error())
		return result, fiber.ErrInternalServerError
	}

	var (
		availableProducts    []*domain.CartList
		nonAvailableProducts []*domain.CartList
	)
	for _, detail := range cart.CartDetails {
		temp := &domain.CartList{
			ID:           detail.ID,
			ProductID:    detail.Product.ID,
			ProductName:  detail.Product.Name,
			ProductPrice: detail.Product.Price,
			Quantity:     detail.Quantity,
		}

		if detail.Product.Stock < 1 {
			nonAvailableProducts = append(nonAvailableProducts, temp)
			continue
		}

		availableProducts = append(availableProducts, temp)
	}

	result = &domain.GetCartListResponse{
		CartID:               cart.ID,
		AvailableProducts:    availableProducts,
		NonAvailableProducts: nonAvailableProducts,
	}

	return result, nil
}

func (c *CartUseCase) Checkout(ctx context.Context, request *domain.CheckoutRequest) (*domain.CheckoutResponse, error) {
	var (
		shippingPrice uint32 = 0 // Assuming the shipping price is free
		result        *domain.CheckoutResponse
	)

	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return result, fiber.ErrBadRequest
	}

	total, _ := c.CartRepo.CountByUserID(request.UserID)
	if total < 1 {
		return result, nil
	}

	cart, err := c.CartRepo.FindByUserID(request.UserID)
	if err != nil {
		log.Errorf("failed to get cart, err: %s", err.Error())
		return result, fiber.ErrInternalServerError
	}

	var (
		totalPrice   uint32
		totalWeight  float32
		orderDetails []*domain.OrderDetail
	)
	for _, detail := range cart.CartDetails {
		if detail.Product.Stock < 1 {
			continue
		}

		temp := &domain.OrderDetail{
			ProductID:      detail.Product.ID,
			ProductName:    detail.Product.Name,
			Quantity:       detail.Quantity,
			Price:          detail.Product.Price,
			Weight:         detail.Product.Weight,
			SubTotalWeight: detail.Product.Weight * float32(detail.Quantity),
			SubTotalPrice:  uint32(detail.Product.Price) * uint32(detail.Quantity),
			CategoryID:     uint64(detail.Product.CategoryID),
		}

		totalWeight += temp.SubTotalWeight
		totalPrice += temp.SubTotalPrice

		orderDetails = append(orderDetails, temp)
	}

	result = &domain.CheckoutResponse{
		CartID:       cart.ID,
		OrderDetails: orderDetails,
		Customer: domain.Customer{
			UserID: cart.User.ID,
		},
		Payment: domain.Payment{
			TotalPrice:   totalPrice,
			TotalPayment: totalPrice + shippingPrice,
		},
		Shipment: domain.Shipment{
			ShippingPrice: shippingPrice,
			TotalWeight:   totalWeight,
			AddressID:     cart.User.Addresses[0].ID,
		},
	}

	return result, nil
}
