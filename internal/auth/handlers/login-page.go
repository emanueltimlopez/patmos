package authHandlers

import (
	"html/template"
	"net/http"
)

var Tmpl *template.Template

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	Tmpl.ExecuteTemplate(w, "login.html", nil)
}
