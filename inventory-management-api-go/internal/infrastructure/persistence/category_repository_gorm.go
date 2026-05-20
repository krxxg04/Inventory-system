package persistence

import (
	"context"
	"errors"

	"inventory-management-api-go/internal/domain/category"
	"inventory-management-api-go/internal/domain/product"
	"inventory-management-api-go/internal/infrastructure/database"

	"gorm.io/gorm"
)

type CategoryRepositoryGorm struct {
	db *gorm.DB
}

func NewCategoryRepositoryGorm(db *gorm.DB) *CategoryRepositoryGorm {
	return &CategoryRepositoryGorm{db: db}
}

func (r *CategoryRepositoryGorm) Create(ctx context.Context, entity *category.Category) error {
	return database.DBFromContext(ctx, r.db).Create(entity).Error
}

func (r *CategoryRepositoryGorm) List(ctx context.Context) ([]category.Category, error) {
	var categories []category.Category
	err := database.DBFromContext(ctx, r.db).
		Order("id ASC").
		Find(&categories).Error
	return categories, err
}

func (r *CategoryRepositoryGorm) GetByID(ctx context.Context, id uint) (*category.Category, error) {
	var entity category.Category
	err := database.DBFromContext(ctx, r.db).First(&entity, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *CategoryRepositoryGorm) Update(ctx context.Context, entity *category.Category) error {
	return database.DBFromContext(ctx, r.db).Save(entity).Error
}

func (r *CategoryRepositoryGorm) Delete(ctx context.Context, id uint) error {
	return database.DBFromContext(ctx, r.db).Delete(&category.Category{}, id).Error
}

func (r *CategoryRepositoryGorm) HasProducts(ctx context.Context, id uint) (bool, error) {
	var count int64
	err := database.DBFromContext(ctx, r.db).
		Model(&product.Product{}).
		Where("category_id = ?", id).
		Count(&count).Error
	return count > 0, err
}
