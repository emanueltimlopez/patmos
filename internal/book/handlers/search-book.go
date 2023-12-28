package book

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	embed "github.com/emanueltimlopez/books-motivation"
	"github.com/emanueltimlopez/books-motivation/internal/book"
	supa "github.com/nedpals/supabase-go"
)

func SearchBookHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	r.ParseForm()
	search := r.Form.Get("search")

	response, err := http.Get(fmt.Sprintf("https://openlibrary.org/search.json?q=%s", url.QueryEscape(search)))

	if err != nil {
		fmt.Print(err.Error())
	}

	booksData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var books book.BooksFromAPI
	_err := json.Unmarshal(booksData, &books)
	if err != nil {
		fmt.Println("Error:", _err)
		return
	}

	embed.TmplComponents.ExecuteTemplate(w, "/search-books.html", books.Docs)
}
