package shelf

import (
	"context"
)

type Repository interface {
	GetShelfByID(ctx context.Context, id string) (*BookShelf, error)
	CreateShelf(ctx context.Context, shelf BookShelf) error
	UpdateShelf(ctx context.Context, id string, shelf BookShelf) error
}
