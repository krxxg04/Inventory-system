package inventory

import (
	"time"

	"inventory-management-api-go/internal/domain/product"
)

type MovementType string

const (
	MovementTypeInbound  MovementType = "INBOUND"
	MovementTypeOutbound MovementType = "OUTBOUND"
)

type InventoryMovement struct {
	ID          uint            `gorm:"primaryKey"`
	ProductID   uint            `gorm:"not null;index"`
	Product     product.Product `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;foreignKey:ProductID;references:ID"`
	Type        MovementType    `gorm:"size:20;not null;index"`
	Quantity    int64           `gorm:"not null;check:quantity > 0"`
	Description string          `gorm:"size:255"`
	CreatedAt   time.Time
}
