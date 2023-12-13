package handlers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/emanueltimlopez/books-motivation/internal/bookshelf"
	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	"github.com/emanueltimlopez/books-motivation/internal/user"
)

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	dbClient := supabase.InitClient()
	dbRepository := supabase.NewSupabaseRepository(dbClient)

	userService := user.NewUserService(dbRepository)
	bookshelfService := bookshelf.NewBookShelfService(dbRepository)

	user, err := userService.GetUser(ctx, "b2e4f2f4-2298-4dd4-9d0f-3b57810ac1a5")
	if err != nil {
		fmt.Println(err)
	}

	books, err := bookshelfService.GetBookShelf(ctx, user.ID.String())
	if err != nil {
		fmt.Println(err)
	}

	tmpl := template.Must(template.ParseFiles("./web/templates/books.html"))
	tmpl.Execute(w, books)
}
