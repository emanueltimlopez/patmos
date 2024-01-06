package book

import (
	"net/http"
	"strings"

	book_usecases "github.com/emanueltimlopez/patmos/internal/book/use-cases"
	supa "github.com/nedpals/supabase-go"
)

type BookFromView struct {
	Title  string
	Author string
	Pages  int
	Image  string
	Isbn   string
}

func SetBookHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	params := book_usecases.GetParamsFormBook(r)

	finalAuthor := strings.Join(params.Author, ",")
	finalIsbn := strings.Join(params.Isbn, ",")
	imageURL := "https://covers.openlibrary.org/b/olid/" + params.Image + "-M.jpg"

	TmplComponents.ExecuteTemplate(w, "book-form.html", BookFromView{
		Title:  params.Title,
		Author: finalAuthor,
		Pages:  params.Pages,
		Image:  imageURL,
		Isbn:   finalIsbn,
	})
}
