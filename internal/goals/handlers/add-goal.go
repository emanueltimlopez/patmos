package goals

import (
	"html/template"
	"net/http"

	supa "github.com/nedpals/supabase-go"
)

func AddGoalHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	tmpl := template.Must(template.ParseFiles("./web/templates/add-goal.html"))
	tmpl.Execute(w, nil)
}
