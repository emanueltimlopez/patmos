package book

import (
	"html/template"
	"net/http"

	supa "github.com/nedpals/supabase-go"
)

func AddBookHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	tmpl := template.Must(template.ParseFiles("./web/templates/add-book.html"))
	tmpl.Execute(w, nil)
}
