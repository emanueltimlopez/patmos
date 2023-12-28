package authHandlers

import (
	"net/http"
)

func SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	Tmpl.ExecuteTemplate(w, "signup.html", nil)
}
