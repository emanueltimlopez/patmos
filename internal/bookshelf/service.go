package bookshelf

import (
	"context"

	"github.com/emanueltimlopez/patmos/internal/book"
)

type Service struct {
	repo Repository
}

func NewBookShelfService(bsr Repository) *Service {
	return &Service{
		repo: bsr,
	}
}

func (bss *Service) GetBookShelf(ctx context.Context, id string) ([]book.Book, error) {
	bs, err := bss.repo.GetUserBooks(ctx, id)

	newBs := []book.Book{}
	for _, value := range bs {
		newBs = append(newBs, value.Book)
	}

	return newBs, err
}
