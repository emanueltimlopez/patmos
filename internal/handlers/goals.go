package handlers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/emanueltimlopez/books-motivation/internal/goals"
	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
)

func GoalsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	dbClient := supabase.InitClient()
	dbRepository := supabase.NewSupabaseRepository(dbClient)

	goalsService := goals.NewGoalService(dbRepository)
	goals, err := goalsService.GetGoals(ctx, "b2e4f2f4-2298-4dd4-9d0f-3b57810ac1a5")
	if err != nil {
		fmt.Println(err)
	}

	tmpl := template.Must(template.ParseFiles("./web/templates/goals.html"))
	tmpl.Execute(w, goals)
}
