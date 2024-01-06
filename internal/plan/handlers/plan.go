package plan

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/emanueltimlopez/patmos/internal/plan"
	planUseCases "github.com/emanueltimlopez/patmos/internal/plan/use-cases"
	"github.com/emanueltimlopez/patmos/internal/platform/supabase"
	"github.com/emanueltimlopez/patmos/internal/user"
	supa "github.com/nedpals/supabase-go"
)

var Tmpl *template.Template

type PlanView struct {
	Plan plan.Plan
	Left int
}

func PlanHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	ctx := context.Background()
	dbClient := supabase.InitClient()
	dbRepository := supabase.NewSupabaseRepository(dbClient)

	userService := user.NewUserService(dbRepository)
	user, err := userService.GetUser(ctx, userSupa.ID)
	if err != nil {
		fmt.Println(err)
	}

	booksLeft := planUseCases.BooksLeft(user)

	Tmpl.ExecuteTemplate(w, "plan.html", PlanView{
		Plan: user.Plan,
		Left: booksLeft,
	})
}
