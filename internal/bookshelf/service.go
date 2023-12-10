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

func (bss *Service) GetBookShelf(ctx context.Context, id string) (*BookShelf, error) {
	bs, err := bss.repo.GetUserBooks(ctx, id)
	bookshelf := &BookShelf{Books: bs}
	return bookshelf, err
}

func (bss *Service) AddBook(ctx context.Context, id string, book book.Book) error {
	err := bss.repo.AddBook(ctx, id, book)
	return err
}
