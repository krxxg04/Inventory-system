package product

import "context"

type Repository interface {
	Create(ctx context.Context, product *Product) error
	List(ctx context.Context) ([]Product, error)
	GetByID(ctx context.Context, id uint) (*Product, error)
	GetByIDForUpdate(ctx context.Context, id uint) (*Product, error)
	GetBySKU(ctx context.Context, sku string) (*Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id uint) error
}
