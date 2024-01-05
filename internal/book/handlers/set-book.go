package book

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/emanueltimlopez/books-motivation/internal/book"
	supa "github.com/nedpals/supabase-go"
)

func SetBookHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	r.ParseForm()
	title := r.Form.Get("title")
	author := r.Form.Get("author")
	isbn := r.Form.Get("isbn")
	pages, _err := strconv.Atoi(r.Form.Get("pages"))
	if _err != nil {
		fmt.Println(_err)
	}

	image := r.Form.Get("cover")
	finalAuthor := strings.ReplaceAll(strings.ReplaceAll(author, "[", ""), "]", "")
	finalIsbn := strings.ReplaceAll(strings.ReplaceAll(isbn, "[", ""), "]", "")

	TmplComponents.ExecuteTemplate(w, "book-form.html", book.Book{
		Title:  title,
		Author: finalAuthor,
		Pages:  pages,
		Image:  image,
		Isbn:   finalIsbn,
	})
}
