package book

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	book_usecases "github.com/emanueltimlopez/patmos/internal/book/use-cases"
	supa "github.com/nedpals/supabase-go"
)

var TmplComponents *template.Template

type BookView struct {
	Title  string
	Pages  int
	Author string
	Image  string
	Isbn   string
}

type SearchBookView struct {
	Books []BookView
}

func SearchBookHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	params := book_usecases.GetParamsSearchBook(r)

	books, err := book_usecases.SearchBook(params.Text)
	if err != nil {
		fmt.Println(err)
	}

	newBooks := []BookView{}
	for _, value := range books {
		newBooks = append(newBooks, BookView{
			Title:  value.Title,
			Pages:  value.Pages,
			Author: strings.Join(value.Author, ","),
			Image:  value.Cover,
			Isbn:   strings.Join(value.Isbn, ","),
		})
	}

	TmplComponents.ExecuteTemplate(w, "search-books.html", SearchBookView{Books: newBooks})
}
