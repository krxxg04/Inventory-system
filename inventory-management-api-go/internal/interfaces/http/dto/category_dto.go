package dto

import "inventory-management-api-go/internal/domain/category"

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type CategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func ToCategoryResponse(entity *category.Category) CategoryResponse {
	return CategoryResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt.UTC().Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   entity.UpdatedAt.UTC().Format("2006-01-02T15:04:05Z"),
	}
}

func ToCategoryResponseList(entities []category.Category) []CategoryResponse {
	items := make([]CategoryResponse, 0, len(entities))
	for _, entity := range entities {
		entityCopy := entity
		items = append(items, ToCategoryResponse(&entityCopy))
	}
	return items
}
