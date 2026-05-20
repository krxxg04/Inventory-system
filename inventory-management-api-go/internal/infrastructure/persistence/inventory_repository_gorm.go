package persistence

import (
	"context"

	"inventory-management-api-go/internal/domain/inventory"
	"inventory-management-api-go/internal/infrastructure/database"

	"gorm.io/gorm"
)

type InventoryRepositoryGorm struct {
	db *gorm.DB
}

func NewInventoryRepositoryGorm(db *gorm.DB) *InventoryRepositoryGorm {
	return &InventoryRepositoryGorm{db: db}
}

func (r *InventoryRepositoryGorm) Create(ctx context.Context, movement *inventory.InventoryMovement) error {
	return database.DBFromContext(ctx, r.db).Create(movement).Error
}

func (r *InventoryRepositoryGorm) ListByProductID(ctx context.Context, productID uint) ([]inventory.InventoryMovement, error) {
	var movements []inventory.InventoryMovement
	err := database.DBFromContext(ctx, r.db).
		Where("product_id = ?", productID).
		Order("created_at DESC, id DESC").
		Find(&movements).Error
	return movements, err
}

func (r *InventoryRepositoryGorm) ExistsByProductID(ctx context.Context, productID uint) (bool, error) {
	var count int64
	err := database.DBFromContext(ctx, r.db).
		Model(&inventory.InventoryMovement{}).
		Where("product_id = ?", productID).
		Count(&count).Error
	return count > 0, err
}
