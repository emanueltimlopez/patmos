package bookshelf

import (
	"context"

	"github.com/emanueltimlopez/books-motivation/internal/book"
)

type Repository interface {
	GetUserBooks(ctx context.Context, id string) ([]*UserBooks, error)
	CreateBook(ctx context.Context, book book.Book, id string) error
}
