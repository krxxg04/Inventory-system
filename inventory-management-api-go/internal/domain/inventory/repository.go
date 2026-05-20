package inventory

import "context"

type Repository interface {
	Create(ctx context.Context, movement *InventoryMovement) error
	ListByProductID(ctx context.Context, productID uint) ([]InventoryMovement, error)
	ExistsByProductID(ctx context.Context, productID uint) (bool, error)
}
