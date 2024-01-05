package bookshelf

import (
	"context"
)

type Repository interface {
	GetUserBooks(ctx context.Context, id string) ([]UserBooksGet, error)
}
