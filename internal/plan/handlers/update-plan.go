package plan

import (
	"context"
	"fmt"
	"net/http"

	"github.com/emanueltimlopez/patmos/internal/plan"
	plan_usecases "github.com/emanueltimlopez/patmos/internal/plan/use-cases"
	"github.com/emanueltimlopez/patmos/internal/platform/supabase"
	"github.com/emanueltimlopez/patmos/internal/user"
	supa "github.com/nedpals/supabase-go"
)

func UpdatePlanHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	ctx := context.Background()
	dbClient := supabase.InitClient()
	dbRepository := supabase.NewSupabaseRepository(dbClient)

	params := plan_usecases.GetParamsPlan(r)

	newPlan := plan.Plan{Minutes: params.Minutes, Sessions: params.Sessions, Words: params.Words}
	userService := user.NewUserService(dbRepository)

	user, err := userService.UpdateUserPlan(ctx, newPlan, userSupa.ID)
	if err != nil {
		fmt.Println("[UpdatePlanHandler]", err)
	}

	booksLeft := plan_usecases.BooksLeft(user)

	Tmpl.ExecuteTemplate(w, "plan.html", PlanView{
		Plan: user.Plan,
		Left: booksLeft,
	})
}
