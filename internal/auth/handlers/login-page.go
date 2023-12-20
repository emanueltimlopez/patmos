package authHandlers

import (
	"html/template"
	"net/http"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/login.html"))
	tmpl.Execute(w, nil)
}
