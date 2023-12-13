package supabase

import (
	"context"
	"fmt"

	"github.com/emanueltimlopez/books-motivation/internal/book"
	"github.com/emanueltimlopez/books-motivation/internal/bookshelf"
	"github.com/emanueltimlopez/books-motivation/internal/user"
	supa "github.com/nedpals/supabase-go"
)

type SupabaseRepository struct {
	client *supa.Client
}

func NewSupabaseRepository(client *supa.Client) *SupabaseRepository {
	return &SupabaseRepository{
		client: client,
	}
}

func (sr *SupabaseRepository) GetBookByID(ctx context.Context, id string) (*book.Book, error) {
	var result *book.Book
	err := sr.client.DB.From("books").Select("*").Single().Eq("id", id).Execute(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

func (sr *SupabaseRepository) CreateBook(ctx context.Context, _book book.Book, userID string) error {
	var result []*book.Book
	err := sr.client.DB.From("books").Insert(_book).Execute(&result)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var resultRelation *bookshelf.UserBooks
	newRelation := bookshelf.UserBooks{UserID: userID, BookID: result[0].ID}
	errRelation := sr.client.DB.From("users_books").Insert(newRelation).Execute(&resultRelation)
	if err != nil {
		fmt.Println(err)

		return errRelation
	}

	return nil
}

func (sr *SupabaseRepository) GetUser(ctx context.Context, id string) (*user.User, error) {
	var result *user.User
	err := sr.client.DB.From("users").Select("*").Single().Eq("id", id).Execute(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

func (sr *SupabaseRepository) GetUserBooks(ctx context.Context, id string) ([]*bookshelf.UserBooks, error) {
	var result []*bookshelf.UserBooks
	err := sr.client.DB.From("users_books").Select("*, books(*)").Eq("user_id", id).Execute(&result)
	if err != nil {
		fmt.Println(err)
	}

	return result, err
}

func (sr *SupabaseRepository) AddBook(ctx context.Context, id string, book book.Book) error {
	return nil
}
