package book

import (
	"fmt"
	"html/template"
	"net/http"

	book_usecases "github.com/emanueltimlopez/books-motivation/internal/book/use-cases"
	supa "github.com/nedpals/supabase-go"
)

var TmplComponents *template.Template

type SearchBookView struct {
	Books []book_usecases.BookFromAPI
}

func SearchBookHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	params := book_usecases.GetParamsSearchBook(r)

	books, err := book_usecases.SearchBook(params.Text)
	if err != nil {
		fmt.Println(err)
	}

	TmplComponents.ExecuteTemplate(w, "search-books.html", SearchBookView{Books: books})
}
