package bookshelf

import (
	"context"

	"github.com/emanueltimlopez/books-motivation/internal/book"
)

type Repository interface {
	GetUserBooks(ctx context.Context, id string) ([]book.Book, error)
	AddBook(ctx context.Context, id string, book book.Book) error
}
