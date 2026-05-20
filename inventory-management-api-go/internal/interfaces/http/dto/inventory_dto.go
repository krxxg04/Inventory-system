package dto

import "inventory-management-api-go/internal/domain/inventory"

type InventoryMovementRequest struct {
	ProductID   uint   `json:"product_id" binding:"required"`
	Quantity    int64  `json:"quantity" binding:"required"`
	Description string `json:"description"`
}

type InventoryMovementResponse struct {
	ID          uint   `json:"id"`
	ProductID   uint   `json:"product_id"`
	Type        string `json:"type"`
	Quantity    int64  `json:"quantity"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

func ToInventoryMovementResponse(entity *inventory.InventoryMovement) InventoryMovementResponse {
	return InventoryMovementResponse{
		ID:          entity.ID,
		ProductID:   entity.ProductID,
		Type:        string(entity.Type),
		Quantity:    entity.Quantity,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt.UTC().Format("2006-01-02T15:04:05Z"),
	}
}

func ToInventoryMovementResponseList(entities []inventory.InventoryMovement) []InventoryMovementResponse {
	items := make([]InventoryMovementResponse, 0, len(entities))
	for _, entity := range entities {
		entityCopy := entity
		items = append(items, ToInventoryMovementResponse(&entityCopy))
	}
	return items
}
