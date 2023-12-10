package routes

import (
	"context"
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
	ctx := context.Background()
	dbClient := supabase.InitClient()
	dbRepository := supabase.NewSupabaseRepository(dbClient)

	userService := user.NewUserService(dbRepository)
	booksService := book.NewBookService(dbRepository)

	activeBook, err := userService.GetUserActiveBook(ctx, "000")
	if err != nil {

	}

	book, err := booksService.GetBookByID(ctx, activeBook)
	if err != nil {

	}

	tmpl := template.Must(template.ParseFiles("./web/templates/index.html"))
	tmpl.Execute(w, book)
}

func todayHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/today.html"))
	tmpl.Execute(w, nil)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/books.html"))
	tmpl.Execute(w, nil)
}
