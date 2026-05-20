package inventory

import "time"

type MovementType string

const (
	MovementTypeInbound  MovementType = "INBOUND"
	MovementTypeOutbound MovementType = "OUTBOUND"
)

type InventoryMovement struct {
	ID          uint         `gorm:"primaryKey"`
	ProductID   uint         `gorm:"not null;index"`
	Type        MovementType `gorm:"size:20;not null;index"`
	Quantity    int64        `gorm:"not null;check:quantity > 0"`
	Description string       `gorm:"size:255"`
	CreatedAt   time.Time
}
