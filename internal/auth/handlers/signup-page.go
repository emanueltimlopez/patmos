package authHandlers

import (
	"net/http"

	embed "github.com/emanueltimlopez/books-motivation"
)

func SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	embed.Tmpl.ExecuteTemplate(w, "signup.html", nil)
}
