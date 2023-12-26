package plan

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	"github.com/emanueltimlopez/books-motivation/internal/user"
	supa "github.com/nedpals/supabase-go"
)

func PlanHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	ctx := context.Background()
	dbClient := supabase.InitClient()
	dbRepository := supabase.NewSupabaseRepository(dbClient)

	userService := user.NewUserService(dbRepository)
	user, err := userService.GetUser(ctx, userSupa.ID)
	if err != nil {
		fmt.Println(err)
	}

	tmpl := template.Must(template.ParseFiles("./web/templates/plan.html"))
	tmpl.Execute(w, user.Plan)
}
