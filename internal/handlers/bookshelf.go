package handlers

import (
	"html/template"
	"net/http"
)

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/books.html"))
	tmpl.Execute(w, nil)
}
