package authHandlers

import (
	"html/template"
	"net/http"
)

func SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/signup.html"))
	tmpl.Execute(w, nil)
}
