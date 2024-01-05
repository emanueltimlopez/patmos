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

func (sr *SupabaseRepository) CreateBook(ctx context.Context, _book book.Book, userID string) (*book.Book, error) {
	var result []*book.Book
	err := sr.client.DB.From("books").Insert(_book).Execute(&result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result[0], nil
}

func (sr *SupabaseRepository) AsociateBook(ctx context.Context, relation book.UserBooksRelation) error {
	var resultRelation *bookshelf.UserBooksGet
	errRelation := sr.client.DB.From("users_books").Insert(relation).Execute(&resultRelation)
	if errRelation != nil {
		fmt.Println(errRelation)
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

func (sr *SupabaseRepository) CreateUser(ctx context.Context, _user user.User) error {
	var result []*user.User
	err := sr.client.DB.From("users").Insert(_user).Execute(&result)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (sr *SupabaseRepository) UpdateUserPlan(ctx context.Context, _user *user.User) (*user.User, error) {
	var result []*user.User
	err := sr.client.DB.From("users").Update(_user).Eq("id", _user.ID).Execute(&result)
	if err != nil {
		fmt.Println("[supabase:UpdateUserPlan:update]", err)
	}
	return result[0], err
}

func (sr *SupabaseRepository) GetUserBooks(ctx context.Context, id string) ([]bookshelf.UserBooksGet, error) {
	var result []bookshelf.UserBooksGet
	err := sr.client.DB.From("users_books").Select("*, books(*)").Eq("user_id", id).Execute(&result)
	if err != nil {
		fmt.Println(err)
	}

	return result, err
}
