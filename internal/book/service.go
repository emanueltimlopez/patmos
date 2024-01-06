package book

import (
	"context"
	"fmt"
	"strings"

	book_usecases "github.com/emanueltimlopez/patmos/internal/book/use-cases"
	"github.com/google/uuid"
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

func (bs *Service) CreateBook(ctx context.Context, params book_usecases.ParamsFormBook, id string) error {

	stringAuthors := "[" + strings.Join(params.Author, ",") + "]"
	stringIsbn := "[" + strings.Join(params.Isbn, ",") + "]"

	newBook := Book{
		Title:  params.Title,
		Author: strings.Split(stringAuthors, ","),
		ID:     uuid.New().String(),
		Pages:  params.Pages,
		Image:  params.Image,
		Isbn:   strings.Split(stringIsbn, ","),
	}

	book, err := bs.repo.CreateBook(ctx, newBook, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	relation := UserBooksRelation{UserID: id, BookID: book.ID}
	_err := bs.repo.AsociateBook(ctx, relation)
	if _err != nil {
		fmt.Println(_err)
	}

	return _err
}
