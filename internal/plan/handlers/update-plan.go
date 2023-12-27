package plan

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	planUseCases "github.com/emanueltimlopez/books-motivation/internal/plan/use-cases"

	"github.com/emanueltimlopez/books-motivation/internal/plan"
	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	"github.com/emanueltimlopez/books-motivation/internal/user"
	supa "github.com/nedpals/supabase-go"
)

func UpdatePlanHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	ctx := context.Background()
	dbClient := supabase.InitClient()
	dbRepository := supabase.NewSupabaseRepository(dbClient)

	r.ParseForm()
	minutes, err := strconv.Atoi(r.Form.Get("minutes"))
	if err != nil {
		fmt.Println(err)
	}

	sessions, _err := strconv.Atoi(r.Form.Get("sessions"))
	if _err != nil {
		fmt.Println(_err)
	}

	words, __err := strconv.Atoi(r.Form.Get("words"))
	if __err != nil {
		fmt.Println(_err)
	}

	newPlan := plan.Plan{Minutes: minutes, Sessions: sessions, Words: words}
	userService := user.NewUserService(dbRepository)

	user, ___err := userService.UpdateUserPlan(ctx, newPlan, userSupa.ID)
	if ___err != nil {
		fmt.Println(___err)
	}

	booksLeft := planUseCases.BooksLeft(user)

	data := map[string]any{
		"Plan": user.Plan,
		"Left": booksLeft,
	}

	tmpl := template.Must(template.ParseFiles("./web/templates/plan.html"))
	tmpl.Execute(w, data)
}
