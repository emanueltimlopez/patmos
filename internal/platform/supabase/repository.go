package supabase

import (
	"context"
	"fmt"

	"github.com/emanueltimlopez/books-motivation/internal/book"
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
	var results *book.Book
	err := sr.client.DB.From("books").Select("*").Eq("id", id).Execute(&results)
	fmt.Println(results)

	return results, err
}

func (sr *SupabaseRepository) CreateBook(ctx context.Context, book *book.Book) error {
	return nil
}

func (sr *SupabaseRepository) GetUser(ctx context.Context, id string) (*user.User, error) {
	var results *user.User
	err := sr.client.DB.From("users").Select("*").Eq("id", id).Execute(&results)
	fmt.Println(results)

	return results, err
}
