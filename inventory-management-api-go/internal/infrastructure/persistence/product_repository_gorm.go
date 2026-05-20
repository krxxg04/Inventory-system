package persistence

import (
	"context"
	"errors"

	"inventory-management-api-go/internal/domain/product"
	"inventory-management-api-go/internal/infrastructure/database"

	"gorm.io/gorm"
)

type ProductRepositoryGorm struct {
	db *gorm.DB
}

func NewProductRepositoryGorm(db *gorm.DB) *ProductRepositoryGorm {
	return &ProductRepositoryGorm{db: db}
}

func (r *ProductRepositoryGorm) Create(ctx context.Context, entity *product.Product) error {
	return database.DBFromContext(ctx, r.db).Create(entity).Error
}

func (r *ProductRepositoryGorm) List(ctx context.Context) ([]product.Product, error) {
	var products []product.Product
	err := database.DBFromContext(ctx, r.db).
		Order("id ASC").
		Find(&products).Error
	return products, err
}

func (r *ProductRepositoryGorm) GetByID(ctx context.Context, id uint) (*product.Product, error) {
	var entity product.Product
	err := database.DBFromContext(ctx, r.db).First(&entity, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *ProductRepositoryGorm) GetByIDForUpdate(ctx context.Context, id uint) (*product.Product, error) {
	var entity product.Product
	tx := database.ForUpdate(database.DBFromContext(ctx, r.db))
	err := tx.First(&entity, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *ProductRepositoryGorm) GetBySKU(ctx context.Context, sku string) (*product.Product, error) {
	var entity product.Product
	err := database.DBFromContext(ctx, r.db).
		Where("sku = ?", sku).
		First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *ProductRepositoryGorm) Update(ctx context.Context, entity *product.Product) error {
	return database.DBFromContext(ctx, r.db).Save(entity).Error
}

func (r *ProductRepositoryGorm) Delete(ctx context.Context, id uint) error {
	return database.DBFromContext(ctx, r.db).Delete(&product.Product{}, id).Error
}
