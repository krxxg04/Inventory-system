package dto

import "inventory-management-api-go/internal/domain/product"

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	SKU         string  `json:"sku" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int64   `json:"stock"`
	CategoryID  uint    `json:"category_id" binding:"required"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	SKU         string  `json:"sku" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	CategoryID  uint    `json:"category_id" binding:"required"`
}

type ProductResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SKU         string  `json:"sku"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
	CategoryID  uint    `json:"category_id"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type ProductStockResponse struct {
	ProductID uint  `json:"product_id"`
	Stock     int64 `json:"stock"`
}

func ToProductResponse(entity *product.Product) ProductResponse {
	return ProductResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		SKU:         entity.SKU,
		Price:       entity.Price,
		Stock:       entity.Stock,
		CategoryID:  entity.CategoryID,
		CreatedAt:   entity.CreatedAt.UTC().Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   entity.UpdatedAt.UTC().Format("2006-01-02T15:04:05Z"),
	}
}

func ToProductResponseList(entities []product.Product) []ProductResponse {
	items := make([]ProductResponse, 0, len(entities))
	for _, entity := range entities {
		entityCopy := entity
		items = append(items, ToProductResponse(&entityCopy))
	}
	return items
}
