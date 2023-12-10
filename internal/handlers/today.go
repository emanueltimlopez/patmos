package handlers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/emanueltimlopez/books-motivation/internal/book"
	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	"github.com/emanueltimlopez/books-motivation/internal/user"
)

func TodayHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	dbClient := supabase.InitClient()
	dbRepository := supabase.NewSupabaseRepository(dbClient)

	userService := user.NewUserService(dbRepository)
	booksService := book.NewBookService(dbRepository)

	user, err := userService.GetUser(ctx, "b2e4f2f4-2298-4dd4-9d0f-3b57810ac1a5")
	if err != nil {
		fmt.Println(err)
	}

	book, err := booksService.GetBookByID(ctx, user.Book)
	if err != nil {
		fmt.Println(err)
	}

	tmpl := template.Must(template.ParseFiles("./web/templates/today.html"))
	tmpl.Execute(w, book)
}
