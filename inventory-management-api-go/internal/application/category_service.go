package application

import (
	"context"
	"strings"

	"inventory-management-api-go/internal/domain/category"
)

type CategoryService struct {
	categoryRepo category.Repository
}

func NewCategoryService(categoryRepo category.Repository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (s *CategoryService) Create(ctx context.Context, name, description string) (*category.Category, error) {
	name = strings.TrimSpace(name)
	description = strings.TrimSpace(description)

	if name == "" {
		return nil, NewValidationError("category name is required")
	}

	entity := &category.Category{
		Name:        name,
		Description: description,
	}

	if err := s.categoryRepo.Create(ctx, entity); err != nil {
		return nil, NewInternalError("failed to create category", err)
	}
	return entity, nil
}

func (s *CategoryService) List(ctx context.Context) ([]category.Category, error) {
	categories, err := s.categoryRepo.List(ctx)
	if err != nil {
		return nil, NewInternalError("failed to list categories", err)
	}
	return categories, nil
}

func (s *CategoryService) GetByID(ctx context.Context, id uint) (*category.Category, error) {
	entity, err := s.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return nil, NewInternalError("failed to get category", err)
	}
	if entity == nil {
		return nil, NewNotFoundError("category not found")
	}
	return entity, nil
}

func (s *CategoryService) Update(ctx context.Context, id uint, name, description string) (*category.Category, error) {
	name = strings.TrimSpace(name)
	description = strings.TrimSpace(description)

	if name == "" {
		return nil, NewValidationError("category name is required")
	}

	entity, err := s.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return nil, NewInternalError("failed to get category", err)
	}
	if entity == nil {
		return nil, NewNotFoundError("category not found")
	}

	entity.Name = name
	entity.Description = description

	if err := s.categoryRepo.Update(ctx, entity); err != nil {
		return nil, NewInternalError("failed to update category", err)
	}
	return entity, nil
}

func (s *CategoryService) Delete(ctx context.Context, id uint) error {
	entity, err := s.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return NewInternalError("failed to get category", err)
	}
	if entity == nil {
		return NewNotFoundError("category not found")
	}

	hasProducts, err := s.categoryRepo.HasProducts(ctx, id)
	if err != nil {
		return NewInternalError("failed to validate category deletion", err)
	}
	if hasProducts {
		return NewConflictError("cannot delete category with associated products")
	}

	if err := s.categoryRepo.Delete(ctx, id); err != nil {
		return NewInternalError("failed to delete category", err)
	}
	return nil
}
