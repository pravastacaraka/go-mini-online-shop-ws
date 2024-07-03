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

type ProductUseCase struct {
	DB          *gorm.DB
	Validate    *validator.Validate
	ProductRepo *repository.ProductRepository
}

func NewProductUseCase(
	db *gorm.DB,
	validate *validator.Validate,
	productRepo *repository.ProductRepository,
) *ProductUseCase {
	return &ProductUseCase{
		DB:          db,
		Validate:    validate,
		ProductRepo: productRepo,
	}
}

func (c *ProductUseCase) Add(ctx context.Context, request *domain.AddProductRequest) (*domain.CommonProductResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return nil, fiber.ErrBadRequest
	}

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	product := &model.Product{
		Name:     request.Name,
		Desc:     request.Desc,
		SKU:      request.SKU,
		Stock:    request.Stock,
		Price:    request.Price,
		Pictures: request.Pictures,
		Category: model.Category{ID: request.CategoryID},
	}
	if err := c.ProductRepo.Create(tx, product); err != nil {
		log.Errorf("failed to add product, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Errorf("failed to commit tx, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &domain.CommonProductResponse{ID: product.ID}, nil
}

func (c *ProductUseCase) Get(ctx context.Context, productID uint64) (*domain.GetProductResponse, error) {
	if err := c.Validate.Var(productID, "gt=0"); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return nil, fiber.ErrBadRequest
	}

	product, err := c.ProductRepo.FindByID(productID)
	if err != nil {
		log.Errorf("failed to get all products, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}
	if product == nil || product.ID < 1 {
		log.Errorf("product id %d doesn't exist!")
		return nil, fiber.ErrNotFound
	}

	result := &domain.GetProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Desc:         product.Desc,
		SKU:          product.SKU,
		Stock:        product.Stock,
		Price:        product.Price,
		Pictures:     product.Pictures,
		CategoryID:   product.Category.ID,
		CategoryName: product.Category.Name,
	}

	return result, nil
}

func (c *ProductUseCase) Update(ctx context.Context, request *domain.UpdateProductRequest) (*domain.CommonProductResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return nil, fiber.ErrBadRequest
	}

	product, err := c.ProductRepo.FindByID(request.ID)
	if err != nil {
		log.Errorf("failed to get all products, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}
	if product == nil || product.ID < 1 {
		log.Errorf("product id %d doesn't exist!")
		return nil, fiber.ErrNotFound
	}

	if request.Name != "" {
		product.Name = request.Name
	}

	if request.Desc != "" {
		product.Desc = request.Desc
	}

	if request.SKU != "" {
		product.SKU = request.SKU
	}

	if request.Stock > 0 {
		product.Stock = request.Stock
	}

	if request.Price > 0 {
		product.Price = request.Price
	}

	if request.CategoryID > 0 {
		product.Category.ID = request.CategoryID
	}

	if len(request.Pictures) > 0 {
		product.Pictures = request.Pictures
	}

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.ProductRepo.Update(tx, product); err != nil {
		log.Errorf("failed to add product, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		log.Errorf("failed to commit tx, err: %s", err.Error())
		return nil, fiber.ErrInternalServerError
	}

	return &domain.CommonProductResponse{ID: product.ID}, nil
}

func (c *ProductUseCase) List(ctx context.Context, request *domain.GetProductsRequest) ([]*domain.GetProductResponse, uint16, error) {
	if err := c.Validate.Struct(request); err != nil {
		log.Errorf("bad request, err: %s", err.Error())
		return nil, 0, fiber.ErrBadRequest
	}

	products, total, err := c.ProductRepo.Find(request)
	if err != nil {
		log.Errorf("failed to get all products, err: %s", err.Error())
		return nil, 0, fiber.ErrInternalServerError
	}

	var result []*domain.GetProductResponse
	for _, product := range products {
		temp := &domain.GetProductResponse{
			ID:           product.ID,
			Name:         product.Name,
			Desc:         product.Desc,
			SKU:          product.SKU,
			Stock:        product.Stock,
			Price:        product.Price,
			Pictures:     product.Pictures,
			CategoryID:   product.Category.ID,
			CategoryName: product.Category.Name,
		}
		result = append(result, temp)
	}

	return result, total, nil
}
