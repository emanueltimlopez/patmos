package home

import (
	"net/http"

	embed "github.com/emanueltimlopez/books-motivation"

	supa "github.com/nedpals/supabase-go"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	embed.Tmpl.ExecuteTemplate(w, "index.html", nil)
}
