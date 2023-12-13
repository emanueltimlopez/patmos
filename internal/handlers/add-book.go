package handlers

import (
	"html/template"
	"net/http"
)

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/add-book.html"))
	tmpl.Execute(w, nil)
}
