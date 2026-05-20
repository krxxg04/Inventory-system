package application

import (
	"context"
	"strings"

	"inventory-management-api-go/internal/domain/category"
	"inventory-management-api-go/internal/domain/inventory"
	"inventory-management-api-go/internal/domain/product"
)

type ProductService struct {
	productRepo   product.Repository
	categoryRepo  category.Repository
	inventoryRepo inventory.Repository
}

func NewProductService(
	productRepo product.Repository,
	categoryRepo category.Repository,
	inventoryRepo inventory.Repository,
) *ProductService {
	return &ProductService{
		productRepo:   productRepo,
		categoryRepo:  categoryRepo,
		inventoryRepo: inventoryRepo,
	}
}

func (s *ProductService) Create(
	ctx context.Context,
	name, description, sku string,
	price float64,
	stock int64,
	categoryID uint,
) (*product.Product, error) {
	name = strings.TrimSpace(name)
	description = strings.TrimSpace(description)
	sku = strings.TrimSpace(sku)

	if name == "" {
		return nil, NewValidationError("product name is required")
	}
	if sku == "" {
		return nil, NewValidationError("product SKU is required")
	}
	if price < 0 {
		return nil, NewValidationError("product price must be greater than or equal to zero")
	}
	if stock < 0 {
		return nil, NewValidationError("product stock must be greater than or equal to zero")
	}

	categoryEntity, err := s.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		return nil, NewInternalError("failed to validate category", err)
	}
	if categoryEntity == nil {
		return nil, NewValidationError("category does not exist")
	}

	existingBySKU, err := s.productRepo.GetBySKU(ctx, sku)
	if err != nil {
		return nil, NewInternalError("failed to validate SKU uniqueness", err)
	}
	if existingBySKU != nil {
		return nil, NewConflictError("product SKU already exists")
	}

	entity := &product.Product{
		Name:        name,
		Description: description,
		SKU:         sku,
		Price:       price,
		Stock:       stock,
		CategoryID:  categoryID,
	}

	if err := s.productRepo.Create(ctx, entity); err != nil {
		return nil, NewInternalError("failed to create product", err)
	}
	return entity, nil
}

func (s *ProductService) List(ctx context.Context) ([]product.Product, error) {
	products, err := s.productRepo.List(ctx)
	if err != nil {
		return nil, NewInternalError("failed to list products", err)
	}
	return products, nil
}

func (s *ProductService) GetByID(ctx context.Context, id uint) (*product.Product, error) {
	entity, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, NewInternalError("failed to get product", err)
	}
	if entity == nil {
		return nil, NewNotFoundError("product not found")
	}
	return entity, nil
}

func (s *ProductService) Update(
	ctx context.Context,
	id uint,
	name, description, sku string,
	price float64,
	categoryID uint,
) (*product.Product, error) {
	name = strings.TrimSpace(name)
	description = strings.TrimSpace(description)
	sku = strings.TrimSpace(sku)

	if name == "" {
		return nil, NewValidationError("product name is required")
	}
	if sku == "" {
		return nil, NewValidationError("product SKU is required")
	}
	if price < 0 {
		return nil, NewValidationError("product price must be greater than or equal to zero")
	}

	entity, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, NewInternalError("failed to get product", err)
	}
	if entity == nil {
		return nil, NewNotFoundError("product not found")
	}

	categoryEntity, err := s.categoryRepo.GetByID(ctx, categoryID)
	if err != nil {
		return nil, NewInternalError("failed to validate category", err)
	}
	if categoryEntity == nil {
		return nil, NewValidationError("category does not exist")
	}

	existingBySKU, err := s.productRepo.GetBySKU(ctx, sku)
	if err != nil {
		return nil, NewInternalError("failed to validate SKU uniqueness", err)
	}
	if existingBySKU != nil && existingBySKU.ID != id {
		return nil, NewConflictError("product SKU already exists")
	}

	entity.Name = name
	entity.Description = description
	entity.SKU = sku
	entity.Price = price
	entity.CategoryID = categoryID

	if err := s.productRepo.Update(ctx, entity); err != nil {
		return nil, NewInternalError("failed to update product", err)
	}
	return entity, nil
}

func (s *ProductService) Delete(ctx context.Context, id uint) error {
	entity, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return NewInternalError("failed to get product", err)
	}
	if entity == nil {
		return NewNotFoundError("product not found")
	}

	hasMovements, err := s.inventoryRepo.ExistsByProductID(ctx, id)
	if err != nil {
		return NewInternalError("failed to validate product deletion", err)
	}
	if hasMovements {
		return NewConflictError("cannot delete product with inventory movements")
	}

	if err := s.productRepo.Delete(ctx, id); err != nil {
		return NewInternalError("failed to delete product", err)
	}
	return nil
}

func (s *ProductService) GetStock(ctx context.Context, id uint) (int64, error) {
	entity, err := s.productRepo.GetByID(ctx, id)
	if err != nil {
		return 0, NewInternalError("failed to get product stock", err)
	}
	if entity == nil {
		return 0, NewNotFoundError("product not found")
	}
	return entity.Stock, nil
}
