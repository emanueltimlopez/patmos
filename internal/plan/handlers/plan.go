package plan

import (
	"context"
	"fmt"
	"net/http"

	embed "github.com/emanueltimlopez/books-motivation"
	planUseCases "github.com/emanueltimlopez/books-motivation/internal/plan/use-cases"
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

	booksLeft := planUseCases.BooksLeft(user)

	data := map[string]any{
		"Plan": user.Plan,
		"Left": booksLeft,
	}

	embed.Tmpl.ExecuteTemplate(w, "plan.html", data)
}
