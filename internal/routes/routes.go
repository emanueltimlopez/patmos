package routes

import (
	"net/http"

	"github.com/emanueltimlopez/books-motivation/internal/handlers"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/bookshelf", handlers.BooksHandler)
	mux.HandleFunc("/today", handlers.TodayHandler)
	mux.HandleFunc("/add-book", handlers.AddBookHandler)
	mux.HandleFunc("/create-book", handlers.CreateBookHandler)

	return mux
}
