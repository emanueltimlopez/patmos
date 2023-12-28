package book

import (
	"net/http"

	embed "github.com/emanueltimlopez/books-motivation"
	supa "github.com/nedpals/supabase-go"
)

func AddBookHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	embed.Tmpl.ExecuteTemplate(w, "/add-book.html", nil)
}
