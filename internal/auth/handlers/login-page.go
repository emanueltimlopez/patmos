package authHandlers

import (
	"net/http"

	embed "github.com/emanueltimlopez/books-motivation"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	embed.Tmpl.ExecuteTemplate(w, "login.html", nil)
}
