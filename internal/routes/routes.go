package routes

import (
	"net/http"

	"github.com/emanueltimlopez/books-motivation/internal/auth"
	authHandlers "github.com/emanueltimlopez/books-motivation/internal/auth/handlers"
	book "github.com/emanueltimlopez/books-motivation/internal/book/handlers"
	bookshelf "github.com/emanueltimlopez/books-motivation/internal/bookshelf/handlers"
	goals "github.com/emanueltimlopez/books-motivation/internal/goals/handlers"
	home "github.com/emanueltimlopez/books-motivation/internal/home/handlers"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.Handle("/", auth.NewAuth(home.IndexHandler))
	mux.HandleFunc("/login", authHandlers.LoginPageHandler)
	mux.HandleFunc("/signup", authHandlers.SignupPageHandler)
	mux.HandleFunc("/login-user", authHandlers.LoginHandler)
	mux.HandleFunc("/signup-user", authHandlers.SignupHandler)
	mux.HandleFunc("/logout-user", authHandlers.LogoutHandler)
	mux.Handle("/bookshelf", auth.NewAuth(bookshelf.BooksHandler))
	mux.Handle("/today", auth.NewAuth(home.TodayHandler))
	mux.Handle("/add-book", auth.NewAuth(book.AddBookHandler))
	mux.Handle("/create-book", auth.NewAuth(book.CreateBookHandler))
	mux.Handle("/search-book", auth.NewAuth(book.SearchBookHandler))
	mux.Handle("/goals", auth.NewAuth(goals.GoalsHandler))
	mux.Handle("/add-goal", auth.NewAuth(goals.AddGoalHandler))
	mux.Handle("/create-goal", auth.NewAuth(goals.CreateGoalHandler))

	return mux
}
