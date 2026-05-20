package category

import "context"

type Repository interface {
	Create(ctx context.Context, category *Category) error
	List(ctx context.Context) ([]Category, error)
	GetByID(ctx context.Context, id uint) (*Category, error)
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id uint) error
	HasProducts(ctx context.Context, id uint) (bool, error)
}
