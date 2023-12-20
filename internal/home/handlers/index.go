package home

import (
	"html/template"
	"net/http"

	supa "github.com/nedpals/supabase-go"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	tmpl := template.Must(template.ParseFiles("./web/templates/index.html"))
	tmpl.Execute(w, nil)
}
