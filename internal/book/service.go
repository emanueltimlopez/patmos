package book

import "context"

type Service struct {
	repo Repository
}

func NewBookService(bs Repository) *Service {
	return &Service{
		repo: bs,
	}
}

func (bs *Service) GetBookByID(ctx context.Context, id string) (*Book, error) {
	book, err := bs.repo.GetBookByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (bs *Service) CreateBook(ctx context.Context, b *Book) error {
	err := bs.repo.CreateBook(ctx, b)
	return err
}