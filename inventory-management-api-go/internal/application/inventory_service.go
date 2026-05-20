package application

import (
	"context"
	"strings"

	"inventory-management-api-go/internal/domain/inventory"
	"inventory-management-api-go/internal/domain/product"
)

type InventoryService struct {
	productRepo   product.Repository
	inventoryRepo inventory.Repository
	txManager     TxManager
}

func NewInventoryService(
	productRepo product.Repository,
	inventoryRepo inventory.Repository,
	txManager TxManager,
) *InventoryService {
	return &InventoryService{
		productRepo:   productRepo,
		inventoryRepo: inventoryRepo,
		txManager:     txManager,
	}
}

func (s *InventoryService) RegisterInbound(
	ctx context.Context,
	productID uint,
	quantity int64,
	description string,
) (*inventory.InventoryMovement, error) {
	return s.registerMovement(ctx, productID, inventory.MovementTypeInbound, quantity, description)
}

func (s *InventoryService) RegisterOutbound(
	ctx context.Context,
	productID uint,
	quantity int64,
	description string,
) (*inventory.InventoryMovement, error) {
	return s.registerMovement(ctx, productID, inventory.MovementTypeOutbound, quantity, description)
}

func (s *InventoryService) ListMovementsByProductID(
	ctx context.Context,
	productID uint,
) ([]inventory.InventoryMovement, error) {
	productEntity, err := s.productRepo.GetByID(ctx, productID)
	if err != nil {
		return nil, NewInternalError("failed to get product", err)
	}
	if productEntity == nil {
		return nil, NewNotFoundError("product not found")
	}

	movements, err := s.inventoryRepo.ListByProductID(ctx, productID)
	if err != nil {
		return nil, NewInternalError("failed to list inventory movements", err)
	}
	return movements, nil
}

func (s *InventoryService) registerMovement(
	ctx context.Context,
	productID uint,
	movementType inventory.MovementType,
	quantity int64,
	description string,
) (*inventory.InventoryMovement, error) {
	if quantity <= 0 {
		return nil, NewValidationError("quantity must be greater than zero")
	}

	description = strings.TrimSpace(description)
	var savedMovement *inventory.InventoryMovement

	err := s.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		productEntity, err := s.productRepo.GetByIDForUpdate(txCtx, productID)
		if err != nil {
			return NewInternalError("failed to get product for stock update", err)
		}
		if productEntity == nil {
			return NewNotFoundError("product not found")
		}

		switch movementType {
		case inventory.MovementTypeInbound:
			productEntity.Stock += quantity
		case inventory.MovementTypeOutbound:
			if productEntity.Stock < quantity {
				return NewConflictError("insufficient stock for outbound movement")
			}
			productEntity.Stock -= quantity
		default:
			return NewValidationError("invalid movement type")
		}

		if err := s.productRepo.Update(txCtx, productEntity); err != nil {
			return NewInternalError("failed to update product stock", err)
		}

		movement := &inventory.InventoryMovement{
			ProductID:   productID,
			Type:        movementType,
			Quantity:    quantity,
			Description: description,
		}
		if err := s.inventoryRepo.Create(txCtx, movement); err != nil {
			return NewInternalError("failed to register inventory movement", err)
		}

		savedMovement = movement
		return nil
	})

	if err != nil {
		return nil, err
	}
	return savedMovement, nil
}
