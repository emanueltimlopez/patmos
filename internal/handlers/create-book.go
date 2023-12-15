package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/emanueltimlopez/books-motivation/internal/book"
	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	"github.com/google/uuid"
)

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	dbClient := supabase.InitClient()
	dbRepository := supabase.NewSupabaseRepository(dbClient)

	r.ParseForm()
	title := r.Form.Get("title")
	author := r.Form.Get("author")

	booksService := book.NewBookService(dbRepository)
	err := booksService.CreateBook(ctx, book.Book{
		Title:  title,
		Author: author,
		ID:     uuid.New().String(),
	}, "b2e4f2f4-2298-4dd4-9d0f-3b57810ac1a5")

	if err != nil {
		fmt.Println(err)
	}
}
