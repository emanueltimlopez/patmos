package book

import (
	"context"
	"fmt"
	"net/http"

	"github.com/emanueltimlopez/books-motivation/internal/book"
	book_usecases "github.com/emanueltimlopez/books-motivation/internal/book/use-cases"
	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	supa "github.com/nedpals/supabase-go"
)

func CreateBookHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	ctx := context.Background()
	dbClient := supabase.InitClient()
	dbRepository := supabase.NewSupabaseRepository(dbClient)

	params := book_usecases.GetParamsFormBook(r)

	booksService := book.NewBookService(dbRepository)
	err := booksService.CreateBook(ctx, params, userSupa.ID)

	if err != nil {
		fmt.Println(err)
	}
}
