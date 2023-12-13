package book

import (
	"context"
)

type Repository interface {
	GetBookByID(ctx context.Context, id string) (*Book, error)
	CreateBook(ctx context.Context, book Book, id string) error
}
