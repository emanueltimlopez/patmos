package book

import (
	"context"
)

type Service struct {
	repo Repository
}

func NewBookService(bs Repository) *Service {
	return &Service{
		repo: bs,
	}
}

func (bs *Service) GetBook(ctx context.Context, id string) (*Book, error) {
	book, err := bs.repo.GetBookByID(ctx, id)
	return book, err
}

func (bs *Service) CreateBook(ctx context.Context, b Book, id string) error {
	err := bs.repo.CreateBook(ctx, b, id)
	return err
}
