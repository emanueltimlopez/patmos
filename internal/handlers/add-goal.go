package handlers

import (
	"html/template"
	"net/http"
)

func AddGoalHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/templates/add-goal.html"))
	tmpl.Execute(w, nil)
}
