package book

import (
	"html/template"
	"net/http"

	supa "github.com/nedpals/supabase-go"
)

var Tmpl *template.Template

func AddBookHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	Tmpl.ExecuteTemplate(w, "add-book.html", nil)
}
