package bookshelf

import (
	"context"

	"github.com/emanueltimlopez/books-motivation/internal/book"
)

type Service struct {
	repo Repository
}

func NewBookShelfService(bsr Repository) *Service {
	return &Service{
		repo: bsr,
	}
}

func (bss *Service) GetBookShelf(ctx context.Context, id string) ([]*UserBooksGet, error) {
	bs, err := bss.repo.GetUserBooks(ctx, id)
	return bs, err
}

func (bss *Service) CreateBook(ctx context.Context, book book.Book, id string) error {
	err := bss.repo.CreateBook(ctx, book, id)
	return err
}
