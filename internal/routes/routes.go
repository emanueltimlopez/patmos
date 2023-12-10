package routes

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/emanueltimlopez/books-motivation/internal/book"
	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	"github.com/emanueltimlopez/books-motivation/internal/user"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/bookshelf", booksHandler)
	mux.HandleFunc("/today", todayHandler)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/index.html"))
	tmpl.Execute(w, nil)
}

func todayHandler(w http.ResponseWriter, r *http.Request) {
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

func booksHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/books.html"))
	tmpl.Execute(w, nil)
}
