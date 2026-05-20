package database

import (
	"inventory-management-api-go/internal/domain/category"
	"inventory-management-api-go/internal/domain/inventory"
	"inventory-management-api-go/internal/domain/product"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&category.Category{},
		&product.Product{},
		&inventory.InventoryMovement{},
	)
}
